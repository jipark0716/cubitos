package entity

import (
	"fmt"
	"games/cubitos/model"
	baseEvent "games/shared/event"
	"games/shared/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type PersonalBoardEntity struct {
	DiceEventChannel   chan *baseEvent.DiceEvent[model.DiceResult]
	RequestIdGenerator <-chan uint64
	dice               *DiceEntity
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
	p.dice.Draw(screen)
}

func NewPersonalBoardEntity(requestIdGenerator <-chan uint64) *PersonalBoardEntity {
	diceEventChannel := make(chan *baseEvent.DiceEvent[model.DiceResult], 64)
	return &PersonalBoardEntity{
		RequestIdGenerator: requestIdGenerator,
		dice:               NewDefaultDiceEntity(diceEventChannel),
		DiceEventChannel:   diceEventChannel,
	}
}
