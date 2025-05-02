package cubitos

import (
	"games/game/entity"
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
	frameCount int
	diceValue  int
	rolling    bool
	entities   []entity.Entity
}

func NewGame() *Game {
	return &Game{
		entities: []entity.Entity{
			entity.NewDefaultDiceEntity(),
		},
	}
}

func (g *Game) Run() {
	ebiten.SetWindowSize(screenWidth*screenRate, screenHeight*screenRate)
	ebiten.SetWindowTitle("2D Dice Animation Example")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	for _, et := range g.entities {
		et.Update()
	}
	return nil
	//if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !g.rolling {
	//	g.rolling = true
	//	g.frameCount = 0
	//}
	//
	//if g.rolling {
	//	g.diceValue = rand.Intn(6) + 1
	//	g.frameCount++
	//	if g.frameCount > rollDuration {
	//		g.rolling = false
	//	}
	//}
	//
	//return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x22, G: 0x22, B: 0x22, A: 0xff})
	for _, et := range g.entities {
		et.Draw(screen)
	}
	//text.Draw(screen, fmt.Sprintf("%d", g.diceValue),
	//	basicfont.Face7x13, screenWidth/2, screenHeight/2, color.White)
	//
	//if g.rolling {
	//	text.Draw(screen, "Rolling...", basicfont.Face7x13, 10, screenHeight-20, color.RGBA{R: 200, G: 200, B: 255, A: 255})
	//} else {
	//	text.Draw(screen, "Click to Roll", basicfont.Face7x13, 10, screenHeight-20, color.RGBA{R: 200, G: 255, B: 200, A: 255})
	//}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / screenRate, outsideHeight / screenRate
}
