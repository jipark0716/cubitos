package cubitos

import (
	"fmt"
	"games/cubitos/entity"
	baseEntity "games/shared/entity"
	baseEvent "games/shared/event"
	"games/shared/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

const (
	screenWidth  = 800
	screenHeight = 600
	screenRate   = 1
)

type Game struct {
	RequestIdGenerator <-chan uint64
	frameCount         int
	diceValue          int
	rolling            bool
	entities           []baseEntity.Entity
	dice               *entity.DefaultDiceEntity
	DiceEventChannel   chan *baseEvent.DiceEvent[*int]
}

func NewGame() *Game {
	diceEventChannel := make(chan *baseEvent.DiceEvent[*int], 64)
	game := &Game{
		DiceEventChannel:   diceEventChannel,
		RequestIdGenerator: util.Increment[uint64](),
		dice:               entity.NewDefaultDiceEntity(diceEventChannel),
	}

	game.EntityReCache()

	return game
}

func (g *Game) EntityReCache() {
	g.entities = []baseEntity.Entity{g.dice}
}

func (g *Game) Run() {
	ebiten.SetWindowSize(screenWidth*screenRate, screenHeight*screenRate)
	ebiten.SetWindowTitle("2D Dice Animation Example")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && g.dice.RollAble() {
		g.dice.Roll(<-g.RequestIdGenerator)
	}

eventHandle:
	for {
		select {
		case evt := <-g.DiceEventChannel:
			fmt.Printf("Dice Event: %v\n", evt)
		default:
			break eventHandle
		}
	}

	for _, et := range g.entities {
		et.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff})
	for _, et := range g.entities {
		et.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / screenRate, outsideHeight / screenRate
}
