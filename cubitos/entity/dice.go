package entity

import (
	"games/cubitos/model"
	baseEntity "games/shared/entity"
	"games/shared/event"
	"games/shared/util"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

type DiceEntity struct {
	*baseEntity.DiceEntity[model.DiceResult]
}

var diceBackground *baseEntity.Drawable
var diceImages []*baseEntity.Drawable

const (
	diceWidth        = float32(100)
	diceHeight       = float32(100)
	diceRadio        = float32(10)
	diceEntityBuffer = 10
)

func init() {
	diceBackground = baseEntity.NewDrawable(
		util.DrawRadiosRect(diceWidth, diceHeight, diceRadio, color.RGBA{R: 180, G: 180, B: 180, A: 255}),
	)

	empty := baseEntity.NewDrawable(ebiten.NewImage(int(diceWidth), int(diceHeight)))

	diceImages = []*baseEntity.Drawable{
		baseEntity.NewDrawable(ebiten.NewImage(int(diceWidth), int(diceHeight))).
			DrawImage(
				util.DrawText("1", basicfont.Face7x13, color.RGBA{R: 255, G: 255, B: 255, A: 255}),
				baseEntity.NewDrawOptions().
					SetScale(4, 4).
					SetPosition(50, 50),
			).
			DrawImage(
				util.DrawCircleOutline(40, 4, color.RGBA{R: 255, G: 255, B: 255, A: 255}),
				baseEntity.NewDrawOptions().
					SetPosition(50, 50),
			),
		empty,
		empty,
		empty,
		empty,
		empty,
	}
}

func NewDefaultDiceEntity(diceEventChannel chan *event.DiceEvent[model.DiceResult]) *DiceEntity {
	return &DiceEntity{
		DiceEntity: &baseEntity.DiceEntity[model.DiceResult]{
			DiceModel:        model.NewDefaultDice(),
			DiceEventChannel: diceEventChannel,
			Background:       diceBackground,
			Images:           diceImages,
			Frame: baseEntity.NewDrawable(
				ebiten.NewImage(int(diceWidth+diceEntityBuffer), int(diceHeight+diceEntityBuffer)),
			),
		},
	}
}
