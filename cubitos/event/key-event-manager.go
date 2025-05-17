package event

import (
	"games/shared/event"
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyEventManager struct {
	RollObserver *event.KeyEventObserver
}

func NewKeyEventManager() *KeyEventManager {
	return &KeyEventManager{
		RollObserver: event.NewKeyEventObserver(ebiten.KeySpace),
	}
}

func (k *KeyEventManager) Update() {
	k.RollObserver.Update()
}
