package entity

import (
	"fmt"
	dice2 "games/cubitos/entity/dice"
	"games/cubitos/model"
	baseEntity "games/shared/entity"
	baseEvent "games/shared/event"
	"games/shared/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type PersonalBoardEntity struct {
	DiceEventChannel   chan *baseEvent.DiceEvent[model.DiceResult]
	RequestIdGenerator <-chan uint64
	Frame              *baseEntity.Drawable
	dice               *dice2.DiceEntity
}

func (p *PersonalBoardEntity) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && p.dice.RollAble() {
		p.dice.Roll(<-p.RequestIdGenerator)
	}
	util.EventHandle(p.DiceEventChannel, func(event *baseEvent.DiceEvent[model.DiceResult]) {
		fmt.Printf("Dice Event: %v\n", event)
	})

	p.dice.Update()
}

func (p *PersonalBoardEntity) Draw(screen *ebiten.Image) {
	frame := p.Frame.CopyWithClear()
	p.dice.Draw(frame.Image)

	screen.DrawImage(p.Frame.Image, &p.Frame.Option)
}

func NewPersonalBoardEntity(requestIdGenerator <-chan uint64) *PersonalBoardEntity {
	diceEventChannel := make(chan *baseEvent.DiceEvent[model.DiceResult], 64)
	return &PersonalBoardEntity{
		RequestIdGenerator: requestIdGenerator,
		dice:               dice2.NewDefaultDiceEntity(diceEventChannel),
		DiceEventChannel:   diceEventChannel,
		Frame:              baseEntity.NewDrawable(ebiten.NewImage(1000, 1000)),
	}
}
