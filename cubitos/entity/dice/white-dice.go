package dice

import (
	"games/cubitos/assets"
	"games/cubitos/model/dice"
	baseEntity "games/shared/entity"
	"games/shared/event"
)

func init() {
	assets.GetFactory().InitGetterAsset(assets.AssetDiceResultWhite, "assets/cubitos/dice-white.png", 5)
}

func NewWhiteDiceEntity(diceEventChannel chan *event.DiceEvent[dice.Result]) *Entity {
	diceDrawOption := baseEntity.NewDrawOptions().SetColorMask(86, 105, 144, 255)

	return &Entity{
		DiceEntity: &baseEntity.DiceEntity[dice.Result]{
			DiceModel:        dice.NewWhiteDice(),
			DiceEventChannel: diceEventChannel,
			Background: assets.GetFactory().
				Get(assets.AssetDiceBackground).
				Copy().
				Translate(
					baseEntity.NewDrawOptions().
						SetColorMask(255, 255, 255, 255),
				),
			Images: []*baseEntity.Drawable{
				assets.GetFactory().Get(assets.AssetDiceResultWhite).Copy().Translate(diceDrawOption),
				assets.GetFactory().Get(assets.AssetDiceResultMove).Copy().Translate(diceDrawOption),
				assets.GetFactory().Get(assets.AssetDiceResultMove).Copy().Translate(diceDrawOption),
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
