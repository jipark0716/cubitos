package entity

import (
	"games/shared/event"
	"games/shared/model"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"math"
)

type Dice interface {
	RollFrame() int
}

type DiceEntity[T any] struct {
	model.DiceModel[T]
	theta               float64
	DiceEventChannel    chan *event.DiceEvent[T]
	frameCount          int
	RequestedFrameCount int
	RollRequestId       uint64
	Frame               *Drawable
	Background          *Drawable
	Images              []*Drawable
}

func (d *DiceEntity[T]) RollFrame() int {
	return 10
}

func (d *DiceEntity[T]) StartRoll() {
	d.SetStatus(model.DiceStatusRolling)
	d.frameCount = 0
}

func (d *DiceEntity[T]) EndRoll(requestId uint64) {
	d.SetStatus(model.DiceStatusRoll)
	d.RequestedFrameCount = d.frameCount + d.RollFrame()
	d.RollRequestId = requestId
}

func (d *DiceEntity[T]) RollComplete() {
	d.DiceModel.Roll()
	d.DiceModel.SetStatus(model.DiceStatusResult)
	d.DiceEventChannel <- &event.DiceEvent[T]{
		Id:            d.RollRequestId,
		DiceEventType: event.DiceEventTypeResult,
		Result:        d.DiceModel.Result(),
	}
}

func (d *DiceEntity[T]) Update() {
	if d.GetStatus() == model.DiceStatusRolling || d.GetStatus() == model.DiceStatusRoll {

		if d.GetStatus() == model.DiceStatusRoll && d.RequestedFrameCount < d.frameCount {
			d.RollComplete()
		} else {
			d.frameCount++
			d.theta += 1
			if d.frameCount%3 == 0 {
				d.DiceModel.Roll()
			}
		}

	}
}

func (d *DiceEntity[T]) Draw(screen *ebiten.Image, options *DrawOptions) {
	frame := d.Frame.CopyWithClear()
	frame.Draw(d.Background.Copy())

	if d.GetStatus() == model.DiceStatusReady {
		frame.Draw(d.Images[0])
	} else {
		frame.Draw(d.Images[d.DiceModel.ResultScreen()])
	}

	frame.SetCenterAnchor()
	if d.GetStatus() == model.DiceStatusRolling || d.GetStatus() == model.DiceStatusRoll {
		frame.Option.GeoM.Rotate(math.Sin(d.theta) * 0.1) // 좌우로 흔들림
	}
	frame.SetStartAnchor()

	frame.Translate(options)
	screen.DrawImage(frame.Image, &frame.Option)
}
