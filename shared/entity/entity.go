package entity

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
)

type Entity interface {
	Update()
	Draw(screen *ebiten.Image)
}

func LoadImages(path string, count int) []*ebiten.Image {
	results := make([]*ebiten.Image, count)
	for i := 0; i < count; i++ {
		path := fmt.Sprintf("%s/%04d.png", path, i+1)
		f, _ := os.Open(path)
		img, _, err := image.Decode(f)

		if err != nil {
			panic(fmt.Sprintf("failed to decode dice animation image: %s", err))
		}

		_ = f.Close()

		results[i] = ebiten.NewImageFromImage(img)
	}

	return results
}
