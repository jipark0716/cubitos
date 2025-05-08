package dice

import (
	"fmt"
	"games/cubitos/assets"
	"games/cubitos/model/dice"
	baseEntity "games/shared/entity"
	"games/shared/util"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"
	"image/png"
	"os"
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
				util.DrawText("1", basicfont.Face7x13),
				baseEntity.NewDrawOptions().
					SetScale(4, 4).
					SetPosition(float64(defaultDiceWidth/2), float64(defaultDiceWidth/2)),
			).
			DrawImage(
				util.DrawCircleOutline(40, 4),
				baseEntity.NewDrawOptions().
					SetPosition(float64(defaultDiceWidth/2), float64(defaultDiceWidth/2)),
			)
	})

	assets.GetFactory().InitGetterImage(assets.AssetDiceBackground, func() *ebiten.Image {
		return util.DrawRadiosRect(defaultDiceWidth, defaultDiceHeight, defaultDiceRadio)
	})

	assets.GetFactory().InitGetterImage(assets.AssetDefaultDiceFrame, func() *ebiten.Image {
		return ebiten.NewImage(int(defaultDiceWidth+defaultDiceEntityBuffer), int(defaultDiceHeight+defaultDiceEntityBuffer))
	})

	assets.GetFactory().InitGetter(assets.AssetDiceResultMove, func() *baseEntity.Drawable {
		path := "cubitos/assets/move.png"
		f, err := os.Open(path)
		if err != nil {
			panic(fmt.Sprintf("open fail: %s error: %v", err, path))
		}

		defer util.PanicIf(f.Close)

		img, err := png.Decode(f)
		if err != nil {
			panic(fmt.Sprintf("decode fail: %s error: %v", err, path))
		}

		return baseEntity.NewDrawable(
			ebiten.NewImageFromImage(img),
		).Translate(baseEntity.NewDrawOptions().SetPosition(5, 5))
	})
}
