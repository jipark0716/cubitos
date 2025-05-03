package model

import model2 "games/shared/model"

type DefaultDice struct {
	model2.BaseDice[*int]
}

var value = 1

func NewDefaultDice() *DefaultDice {
	return &DefaultDice{
		model2.BaseDice[*int]{
			Values: []*int{nil, &value, nil, nil, nil, nil},
		},
	}
}

func (d *DefaultDice) Roll() {
	d.BaseDice.Roll()
	println(d.Result())
}
