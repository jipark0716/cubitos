package entity

import (
	"games/game/model"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Dice interface {
	RollFrame() int
}

type DiceEntity[T any] struct {
	model.DiceModel[T]
	frameCount   int
	animateIndex int
	diceResult   []*ebiten.Image
	diceAnimate  []*ebiten.Image
}

func (d *DiceEntity[T]) RollFrame() int {
	return 30
}

const (
	DiceFrames   = 12
	AnimateSpeed = 2
)

func (d *DiceEntity[T]) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && d.RollAble() {
		d.SetStatus(model.DiceStatusRoll)
		d.frameCount = 0
	}

	if d.GetStatus() == model.DiceStatusRoll {
		d.frameCount++
		if d.frameCount%AnimateSpeed == 0 {
			d.animateIndex = (d.animateIndex + 1) % DiceFrames
		}

		if d.frameCount > d.RollFrame() {
			d.Roll()
		}
	}
}

func (d *DiceEntity[T]) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	var img *ebiten.Image
	if d.GetStatus() == model.DiceStatusResult {
		img = d.diceResult[d.ResultScreen()]
	} else {
		img = d.diceAnimate[d.animateIndex]
	}

	op.Filter = ebiten.FilterNearest

	screen.DrawImage(img, op)
}
