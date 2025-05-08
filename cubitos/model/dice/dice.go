package dice

import baseModel "games/shared/model"

type Dice struct {
	baseModel.BaseDice[Result]
	DiceType Type
}

type Result uint8

const (
	ResultNone Result = iota
	ResultFlushableCoin1
	ResultFlushableCoin2
	ResultFlushableCoin3
	ResultCoin1
	ResultMove1
)

type Type uint8

const (
	TypeGray Type = iota
	TypeBlack
	TypeStart
	TypeRed
	TypeOrange
	TypeYellow
	TypeGreen
	TypeBlue
	TypePurple
	TypeBrown
	TypeWhite
)

var value = 1

func (d *Dice) Roll() {
	d.BaseDice.Roll()
}

func NewGrayDice() *Dice {
	return &Dice{
		DiceType: TypeGray,
		BaseDice: baseModel.BaseDice[Result]{
			Values: []Result{
				ResultFlushableCoin1,
				ResultNone,
				ResultNone,
				ResultNone,
				ResultNone,
				ResultNone,
			},
		},
	}
}

func NewBlackDice() *Dice {
	return &Dice{
		DiceType: TypeGray,
		BaseDice: baseModel.BaseDice[Result]{
			Values: []Result{
				ResultMove1,
				ResultFlushableCoin1,
				ResultNone,
				ResultNone,
				ResultNone,
				ResultNone,
			},
		},
	}
}
