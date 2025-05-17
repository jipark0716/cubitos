package event

import "github.com/hajimehoshi/ebiten/v2"

type KeyEventObserver struct {
	keys      []ebiten.Key
	IsPressed bool
	onKeyUp   []func()
	onKeyDown []func()
}

func NewKeyEventObserver(keys ...ebiten.Key) *KeyEventObserver {
	return &KeyEventObserver{
		keys:      keys,
		IsPressed: false,
	}
}

func (k *KeyEventObserver) isPressed() bool {
	for _, key := range k.keys {
		if !ebiten.IsKeyPressed(key) {
			return false
		}
	}
	return true
}

func (k *KeyEventObserver) Update() {
	if k.IsPressed != k.isPressed() {
		k.IsPressed = !k.IsPressed

		if k.IsPressed {
			k.KeyDown()
		} else {
			k.KeyUp()
		}
	}
}

func (k *KeyEventObserver) AddKeyUpListener(f func()) {
	k.onKeyUp = append(k.onKeyUp, f)
}

func (k *KeyEventObserver) AddKeyDownListener(f func()) {
	k.onKeyDown = append(k.onKeyDown, f)
}

func (k *KeyEventObserver) KeyDown() {
	for _, handler := range k.onKeyDown {
		handler()
	}
}

func (k *KeyEventObserver) KeyUp() {
	for _, handler := range k.onKeyUp {
		handler()
	}
}
