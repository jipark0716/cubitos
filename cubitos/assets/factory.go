package assets

import (
	"bytes"
	"fmt"
	baseAssets "games/shared/assets"
	baseEntity "games/shared/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
)

type Factory struct {
	Getter map[AssetType]func() *baseEntity.Drawable
	Entity map[AssetType]*baseEntity.Drawable
}

type AssetType uint16

const (
	AssetEmpty AssetType = iota
	AssetDefaultDiceFrame
	AssetDiceResultMove
	AssetDiceResultFlushableCoin1
	AssetDiceResultFlushableCoin2
	AssetDiceResultFlushableCoin3
	AssetDiceResultBrown
	AssetDiceResultWhite
	AssetDiceBackground
)

var FactoryInstance *Factory

func init() {
	FactoryInstance = &Factory{
		Getter: map[AssetType]func() *baseEntity.Drawable{},
		Entity: map[AssetType]*baseEntity.Drawable{},
	}

	FactoryInstance.InitGetterImage(AssetEmpty, func() *ebiten.Image {
		return ebiten.NewImage(1, 1)
	})
}

func GetFactory() *Factory {
	return FactoryInstance
}

func (f *Factory) InitGetterImage(assetType AssetType, getter func() *ebiten.Image) {
	if f.Getter[assetType] == nil {
		f.Getter[assetType] = func() *baseEntity.Drawable {
			return baseEntity.NewDrawable(getter())
		}
	} else {
		panic(fmt.Sprintf("%d is already initialized", assetType))
	}
}

func (f *Factory) InitGetter(assetType AssetType, getter func() *baseEntity.Drawable) {
	if f.Getter[assetType] == nil {
		f.Getter[assetType] = getter
	} else {
		panic(fmt.Sprintf("%d is already initialized", assetType))
	}
}

func (f *Factory) InitGetterAsset(assetType AssetType, path string, padding float64) {
	if f.Getter[assetType] != nil {
		panic(fmt.Sprintf("%d is already initialized", assetType))
	}

	f.Getter[assetType] = func() *baseEntity.Drawable {
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
		).Translate(baseEntity.NewDrawOptions().SetPosition(padding, padding))
	}

}

func (f *Factory) Get(assetType AssetType) *baseEntity.Drawable {
	if f.Entity[assetType] == nil {
		if f.Getter[assetType] == nil {
			panic(fmt.Sprintf("%d undefind getter", assetType))
		}
		f.Entity[assetType] = f.Getter[assetType]()
	}
	return f.Entity[assetType]
}
