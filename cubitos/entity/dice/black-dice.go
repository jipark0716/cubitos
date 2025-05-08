package dice

import (
	"games/cubitos/assets"
	"games/cubitos/model/dice"
	baseEntity "games/shared/entity"
	"games/shared/event"
	"games/shared/util"
	"image/color"
)

func init() {
	assets.GetFactory().InitGetter(assets.AssetBlackDiceBackground, func() *baseEntity.Drawable {
		return baseEntity.NewDrawable(
			util.DrawRadiosRect(defaultDiceWidth, defaultDiceHeight, defaultDiceRadio, color.RGBA{R: 120, G: 120, B: 120, A: 255}),
		)
	})
}

func NewBlackDiceEntity(diceEventChannel chan *event.DiceEvent[dice.Result]) *Entity {
	return &Entity{
		DiceEntity: &baseEntity.DiceEntity[dice.Result]{
			DiceModel:        dice.NewGrayDice(),
			DiceEventChannel: diceEventChannel,
			Background:       assets.GetFactory().Get(assets.AssetBlackDiceBackground),
			Images: []*baseEntity.Drawable{
				assets.GetFactory().Get(assets.AssetDiceResultFlushableCoin1),
				assets.GetFactory().Get(assets.AssetEmpty),
				assets.GetFactory().Get(assets.AssetEmpty),
				assets.GetFactory().Get(assets.AssetEmpty),
				assets.GetFactory().Get(assets.AssetEmpty),
				assets.GetFactory().Get(assets.AssetEmpty),
			},
			Frame: assets.GetFactory().
				Get(assets.AssetDefaultDiceFrame).
				Copy().
				Translate(baseEntity.
					NewDrawOptions().
					SetPosition(10, 10),
				),
		},
	}
}
