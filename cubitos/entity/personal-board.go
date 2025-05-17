package entity

import (
	"fmt"
	"games/cubitos/assets"
	diceEntity "games/cubitos/entity/dice"
	"games/cubitos/event"
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
	ReadyDices         []*diceEntity.Entity
	RollDices          []*diceEntity.Entity
	ActiveDices        []*diceEntity.Entity
	TrashDices         []*diceEntity.Entity
	pendingCount       int
	BoardAsset         *baseEntity.Drawable
}

func init() {
	assets.GetFactory().InitGetterAsset(assets.AssetPersonalBoard, "assets/cubitos/personal-board.png", 0)
}

func (p *PersonalBoardEntity) ReCache() {
	p.dices = util.Merge(p.ReadyDices, p.RollDices, p.ActiveDices, p.TrashDices)
}

func (p *PersonalBoardEntity) StartRoll() {
	if p.pendingCount == 0 {
		for _, dice := range p.dices {
			dice.StartRoll()
		}
	}
}

func (p *PersonalBoardEntity) EndRoll() {
	if p.pendingCount == 0 {
		p.pendingCount += len(p.dices)
		for _, dice := range p.dices {
			dice.EndRoll(<-p.RequestIdGenerator)
		}
	}
}

func (p *PersonalBoardEntity) Update() {
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
	frame.Draw(p.BoardAsset)
	baseEntity.DrawTile(
		frame,
		p.ReadyDices,
		baseEntity.NewTitleDrawOption(3).
			SetGPaddingX(10).
			SetGPaddingY(5).
			SetPadding(60))

	screen.DrawImage(p.Frame.Image, &p.Frame.Option)
}

func NewPersonalBoardEntity(requestIdGenerator <-chan uint64, eventManager *event.KeyEventManager) *PersonalBoardEntity {
	diceEventChannel := make(chan *baseEvent.DiceEvent[diceModel.Result], 64)
	p := &PersonalBoardEntity{
		RequestIdGenerator: requestIdGenerator,
		ReadyDices: []*diceEntity.Entity{
			diceEntity.NewGrayDiceEntity(diceEventChannel),
			diceEntity.NewBlackDiceEntity(diceEventChannel),
			diceEntity.NewBrownDiceEntity(diceEventChannel),
			diceEntity.NewWhiteDiceEntity(diceEventChannel),
		},
		DiceEventChannel: diceEventChannel,
		BoardAsset:       assets.GetFactory().Get(assets.AssetPersonalBoard),
		Frame:            baseEntity.NewDrawable(ebiten.NewImage(1000, 400)),
	}
	p.ReCache()
	eventManager.RollObserver.AddKeyDownListener(p.StartRoll)
	eventManager.RollObserver.AddKeyUpListener(p.EndRoll)

	return p
}
