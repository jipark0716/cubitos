package model

import baseModel "games/shared/model"

type Dice struct {
	baseModel.BaseDice[DiceResult]
}

type DiceResult uint8

const (
	DiceResultNone DiceResult = iota
	DiceResultCone1
	DiceResultCone2
	DiceResultConeMove
)

var value = 1

func NewDefaultDice() *Dice {
	return &Dice{
		baseModel.BaseDice[DiceResult]{
			Values: []DiceResult{
				DiceResultNone,
				DiceResultCone1,
				DiceResultNone,
				DiceResultNone,
				DiceResultNone,
				DiceResultNone,
			},
		},
	}
}

func (d *Dice) Roll() {
	d.BaseDice.Roll()
}
