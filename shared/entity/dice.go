package entity

import (
	"games/shared/event"
	"games/shared/model"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Dice interface {
	RollFrame() int
}

type DiceEntity[T any] struct {
	model.DiceModel[T]
	DiceEventChannel chan *event.DiceEvent[T]
	frameCount       int
	animateIndex     int
	RollRequestId    uint64
	DiceResult       []*ebiten.Image
	DiceAnimate      []*ebiten.Image
}

func (d *DiceEntity[T]) RollFrame() int {
	return 30
}

const (
	DiceFrames   = 12
	AnimateSpeed = 2
)

func (d *DiceEntity[T]) Roll(requestId uint64) {
	d.SetStatus(model.DiceStatusRoll)
	d.frameCount = 0
	d.RollRequestId = requestId
}

func (d *DiceEntity[T]) RollComplete() {
	d.DiceModel.Roll()
	d.DiceEventChannel <- &event.DiceEvent[T]{
		Id:            d.RollRequestId,
		DiceEventType: event.DiceEventTypeResult,
		Result:        d.DiceModel.Result(),
	}
}

func (d *DiceEntity[T]) Update() {
	if d.GetStatus() == model.DiceStatusRoll {
		d.frameCount++
		if d.frameCount%AnimateSpeed == 0 {
			d.animateIndex = (d.animateIndex + 1) % DiceFrames
		}

		if d.frameCount > d.RollFrame() {
			d.RollComplete()
		}
	}
}

func (d *DiceEntity[T]) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	var img *ebiten.Image
	if d.GetStatus() == model.DiceStatusResult {
		img = d.DiceResult[d.ResultScreen()]
	} else {
		img = d.DiceAnimate[d.animateIndex]
	}

	op.Filter = ebiten.FilterNearest

	screen.DrawImage(img, op)
}
