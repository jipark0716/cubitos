package entity

import (
	"games/cubitos/model"
	entity2 "games/shared/entity"
	"games/shared/event"
	"github.com/hajimehoshi/ebiten/v2"
)

var diceAnimate []*ebiten.Image
var diceResult []*ebiten.Image

func init() {
	diceAnimate = entity2.LoadImages("cubitos/assets/default_dice_animation", entity2.DiceFrames)
	diceResult = entity2.LoadImages("cubitos/assets/default_dice_result", 6)
}

type DefaultDiceEntity struct {
	*entity2.DiceEntity[*int]
}

func NewDefaultDiceEntity(diceEventChannel chan *event.DiceEvent[*int]) *DefaultDiceEntity {
	return &DefaultDiceEntity{
		DiceEntity: &entity2.DiceEntity[*int]{
			DiceModel:        model.NewDefaultDice(),
			DiceAnimate:      diceAnimate,
			DiceResult:       diceResult,
			DiceEventChannel: diceEventChannel,
		},
	}
}
