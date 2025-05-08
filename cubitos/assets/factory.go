package assets

import (
	"fmt"
	baseEntity "games/shared/entity"
	"github.com/hajimehoshi/ebiten/v2"
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

func (f *Factory) Get(assetType AssetType) *baseEntity.Drawable {
	if f.Entity[assetType] == nil {
		if f.Getter[assetType] == nil {
			panic(fmt.Sprintf("%d undefind getter", assetType))
		}
		f.Entity[assetType] = f.Getter[assetType]()
	}
	return f.Entity[assetType]
}
