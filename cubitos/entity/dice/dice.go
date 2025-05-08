package dice

import (
	"games/cubitos/assets"
	"games/cubitos/model/dice"
	baseEntity "games/shared/entity"
	"games/shared/util"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

type Entity struct {
	*baseEntity.DiceEntity[dice.Result]
}

const (
	defaultDiceWidth        = float32(100)
	defaultDiceHeight       = float32(100)
	defaultDiceRadio        = float32(10)
	defaultDiceEntityBuffer = 10
)

func init() {
	assets.GetFactory().InitGetter(assets.AssetDiceResultFlushableCoin1, func() *baseEntity.Drawable {
		return baseEntity.NewDrawable(ebiten.NewImage(int(defaultDiceWidth), int(defaultDiceHeight))).
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
			)
	})

	assets.GetFactory().InitGetter(assets.AssetDefaultDiceFrame, func() *baseEntity.Drawable {
		return baseEntity.NewDrawable(ebiten.NewImage(int(defaultDiceWidth+defaultDiceEntityBuffer), int(defaultDiceHeight+defaultDiceEntityBuffer)))
	})
}
