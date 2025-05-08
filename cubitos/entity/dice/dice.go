package dice

import (
	"bytes"
	"fmt"
	"games/cubitos/assets"
	"games/cubitos/model/dice"
	baseAssets "games/shared/assets"
	baseEntity "games/shared/entity"
	"games/shared/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
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
		return util.DrawCircleText("1", defaultDiceWidth, defaultDiceHeight)

	})

	assets.GetFactory().InitGetter(assets.AssetDiceResultFlushableCoin2, func() *baseEntity.Drawable {
		return util.DrawCircleText("2", defaultDiceWidth, defaultDiceHeight)
	})

	assets.GetFactory().InitGetter(assets.AssetDiceResultFlushableCoin3, func() *baseEntity.Drawable {
		return util.DrawCircleText("3", defaultDiceWidth, defaultDiceHeight)
	})

	assets.GetFactory().InitGetterImage(assets.AssetDiceBackground, func() *ebiten.Image {
		return util.DrawRadiosRect(defaultDiceWidth, defaultDiceHeight, defaultDiceRadio)
	})

	assets.GetFactory().InitGetterImage(assets.AssetDefaultDiceFrame, func() *ebiten.Image {
		return ebiten.NewImage(int(defaultDiceWidth+defaultDiceEntityBuffer), int(defaultDiceHeight+defaultDiceEntityBuffer))
	})

	assets.GetFactory().InitGetter(assets.AssetDiceResultMove, func() *baseEntity.Drawable {
		path := "assets/cubitos/dice-move.png"
		data, err := baseAssets.GetLoader().Load(path)
		if err != nil {
			panic(fmt.Sprintf("read fail: %s error: %v", path, err))
		}

		img, err := png.Decode(bytes.NewReader(data))
		if err != nil {
			panic(fmt.Sprintf("decode fail: %s error: %v", path, err))
		}

		return baseEntity.NewDrawable(
			ebiten.NewImageFromImage(img),
		).Translate(baseEntity.NewDrawOptions().SetPosition(5, 5))

	})
}
