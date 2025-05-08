package dice

import (
	"games/cubitos/assets"
	"games/cubitos/model/dice"
	baseEntity "games/shared/entity"
	"games/shared/event"
)

func init() {
	assets.GetFactory().InitGetterAsset(assets.AssetDiceResultBrown, "assets/cubitos/dice-brown.png", 5)
}

func NewBrownDiceEntity(diceEventChannel chan *event.DiceEvent[dice.Result]) *Entity {
	return &Entity{
		DiceEntity: &baseEntity.DiceEntity[dice.Result]{
			DiceModel:        dice.NewBrownDice(),
			DiceEventChannel: diceEventChannel,
			Background: assets.GetFactory().
				Get(assets.AssetDiceBackground).
				Copy().
				Translate(
					baseEntity.NewDrawOptions().
						SetColorMask(147, 130, 103, 255),
				),
			Images: []*baseEntity.Drawable{
				assets.GetFactory().Get(assets.AssetDiceResultBrown),
				assets.GetFactory().Get(assets.AssetDiceResultFlushableCoin3),
				assets.GetFactory().Get(assets.AssetEmpty),
				assets.GetFactory().Get(assets.AssetEmpty),
				assets.GetFactory().Get(assets.AssetEmpty),
				assets.GetFactory().Get(assets.AssetEmpty),
			},
			Frame: assets.GetFactory().
				Get(assets.AssetDefaultDiceFrame).
				Copy().
				Translate(
					baseEntity.
						NewDrawOptions().
						SetPosition(10, 10),
				),
		},
	}
}
