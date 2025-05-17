package cubitos

import (
	"games/cubitos/entity"
	"games/cubitos/event"
	baseEntity "games/shared/entity"
	"games/shared/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

const (
	screenWidth  = 1200
	screenHeight = 800
	screenRate   = 1
)

type Game struct {
	RequestIdGenerator <-chan uint64
	frameCount         int
	diceValue          int
	rolling            bool
	entities           []baseEntity.Entity
	KeyEventManager    *event.KeyEventManager
	PersonalBoard      *entity.PersonalBoardEntity
}

func NewGame() *Game {
	requestIdGenerator := util.Increment[uint64]()
	eventManager := event.NewKeyEventManager()
	game := &Game{
		RequestIdGenerator: requestIdGenerator,
		PersonalBoard:      entity.NewPersonalBoardEntity(requestIdGenerator, eventManager),
		KeyEventManager:    eventManager,
	}

	game.EntityReCache()

	return game
}

func (g *Game) EntityReCache() {
	g.entities = []baseEntity.Entity{g.PersonalBoard}
}

func (g *Game) Run() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("2D Dice Animation Example")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	g.KeyEventManager.Update()
	for _, et := range g.entities {
		et.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff})
	for _, et := range g.entities {
		et.Draw(screen, baseEntity.NewDrawOptions())
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth * screenRate, outsideHeight * screenRate
}
