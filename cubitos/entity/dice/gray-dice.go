package dice

import (
	"games/cubitos/assets"
	"games/cubitos/model/dice"
	baseEntity "games/shared/entity"
	"games/shared/event"
)

func NewGrayDiceEntity(diceEventChannel chan *event.DiceEvent[dice.Result]) *Entity {
	return &Entity{
		DiceEntity: &baseEntity.DiceEntity[dice.Result]{
			DiceModel:        dice.NewBlackDice(),
			DiceEventChannel: diceEventChannel,
			Background: assets.GetFactory().
				Get(assets.AssetDiceBackground).
				Copy().
				Translate(
					baseEntity.NewDrawOptions().
						SetColorMask(180, 180, 180, 255),
				),
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
				Translate(
					baseEntity.
						NewDrawOptions().
						SetPosition(10, 10),
				),
		},
	}
}
