package dice

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
	defaultDiceWidth        = float32(100)
	defaultDiceHeight       = float32(100)
	defaultDiceRadio        = float32(10)
	defaultDiceEntityBuffer = 10
)

func init() {
	diceBackground = baseEntity.NewDrawable(
		util.DrawRadiosRect(defaultDiceWidth, defaultDiceHeight, defaultDiceRadio, color.RGBA{R: 180, G: 180, B: 180, A: 255}),
	)

	empty := baseEntity.NewDrawable(ebiten.NewImage(int(defaultDiceWidth), int(defaultDiceHeight)))

	diceImages = []*baseEntity.Drawable{
		baseEntity.NewDrawable(ebiten.NewImage(int(defaultDiceWidth), int(defaultDiceHeight))).
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
				ebiten.NewImage(int(defaultDiceWidth+defaultDiceEntityBuffer), int(defaultDiceHeight+defaultDiceEntityBuffer)),
			).Translate(10, 10),
		},
	}
}
