package entity

import (
	"games/game/model"
	"github.com/hajimehoshi/ebiten/v2"
)

var diceAnimate []*ebiten.Image
var diceResult []*ebiten.Image

func init() {
	diceAnimate = LoadImages("assets/default_dice_animation", DiceFrames)
	diceResult = LoadImages("assets/default_dice_result", 6)
}

type DefaultDiceEntity struct {
	*DiceEntity[*int]
}

func NewDefaultDiceEntity() *DefaultDiceEntity {
	return &DefaultDiceEntity{
		DiceEntity: &DiceEntity[*int]{
			DiceModel:   model.NewDefaultDice(),
			diceAnimate: diceAnimate,
			diceResult:  diceResult,
		},
	}
}
