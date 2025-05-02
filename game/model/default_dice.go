package model

type DefaultDice struct {
	BaseDice[*int]
}

var value = 1

func NewDefaultDice() *DefaultDice {
	return &DefaultDice{
		BaseDice[*int]{
			Values: []*int{nil, &value, nil, nil, nil, nil},
		},
	}
}

func (d *DefaultDice) Roll() {
	d.BaseDice.Roll()
	println(d.Result())
}
