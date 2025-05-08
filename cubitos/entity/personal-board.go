package entity

import (
	"fmt"
	diceEntity "games/cubitos/entity/dice"
	diceModel "games/cubitos/model/dice"
	baseEntity "games/shared/entity"
	baseEvent "games/shared/event"
	"games/shared/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type PersonalBoardEntity struct {
	DiceEventChannel   chan *baseEvent.DiceEvent[diceModel.Result]
	RequestIdGenerator <-chan uint64
	Frame              *baseEntity.Drawable
	dices              []*diceEntity.Entity
	pendingCount       int
}

const (
	DicePadding = 120
)

func (p *PersonalBoardEntity) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && p.pendingCount == 0 {
		p.pendingCount += len(p.dices)
		for _, dice := range p.dices {
			dice.Roll(<-p.RequestIdGenerator)
		}
	}
	util.EventHandle(p.DiceEventChannel, func(event *baseEvent.DiceEvent[diceModel.Result]) {
		p.pendingCount -= 1
		fmt.Printf("Dice Event: %v\n", event)
	})

	for _, dice := range p.dices {
		dice.Update()
	}
}

func (p *PersonalBoardEntity) Draw(screen *ebiten.Image, options *baseEntity.DrawOptions) {
	frame := p.Frame.CopyWithClear()
	frame.Translate(options)
	for i, dice := range p.dices {
		dice.Draw(
			frame.Image,
			baseEntity.
				NewDrawOptions().
				SetPosition(float64(i*DicePadding), 0),
		)
	}

	screen.DrawImage(p.Frame.Image, &p.Frame.Option)
}

func NewPersonalBoardEntity(requestIdGenerator <-chan uint64) *PersonalBoardEntity {
	diceEventChannel := make(chan *baseEvent.DiceEvent[diceModel.Result], 64)
	return &PersonalBoardEntity{
		RequestIdGenerator: requestIdGenerator,
		dices: []*diceEntity.Entity{
			diceEntity.NewGrayDiceEntity(diceEventChannel),
			diceEntity.NewBlackDiceEntity(diceEventChannel),
			diceEntity.NewGrayDiceEntity(diceEventChannel),
			diceEntity.NewBlackDiceEntity(diceEventChannel),
			diceEntity.NewGrayDiceEntity(diceEventChannel),
			diceEntity.NewBlackDiceEntity(diceEventChannel),
			diceEntity.NewGrayDiceEntity(diceEventChannel),
			diceEntity.NewBlackDiceEntity(diceEventChannel),
			diceEntity.NewGrayDiceEntity(diceEventChannel),
			diceEntity.NewBlackDiceEntity(diceEventChannel),
		},
		DiceEventChannel: diceEventChannel,
		Frame:            baseEntity.NewDrawable(ebiten.NewImage(1000, 1000)),
	}
}
