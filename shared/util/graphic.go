package util

import (
	baseEntity "games/shared/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func DrawRadiosRect(w, h, r float32) *ebiten.Image {
	clr := color.White
	image := ebiten.NewImage(int(w), int(h))
	vector.DrawFilledRect(image, r, 0, w-2*r, h, clr, true)
	vector.DrawFilledRect(image, 0, r, w, h-2*r, clr, true)
	vector.DrawFilledCircle(image, r, r, r, clr, true)
	vector.DrawFilledCircle(image, w-r, r, r, clr, true)
	vector.DrawFilledCircle(image, r, h-r, r, clr, true)
	vector.DrawFilledCircle(image, w-r, h-r, r, clr, true)

	return image
}

func DrawText(content string, face *basicfont.Face) *ebiten.Image {
	clr := color.White
	bounds, _ := font.BoundString(face, content)

	w := (bounds.Max.X - bounds.Min.X).Ceil()
	h := (bounds.Max.Y - bounds.Min.Y).Ceil()
	txtImage := ebiten.NewImage(w, h)

	text.Draw(
		txtImage,
		content,
		face,
		0,
		face.Metrics().Ascent.Ceil(),
		clr,
	)

	return txtImage
}

func DrawCircleOutline(radius, thickness float32) *ebiten.Image {
	clr := color.White
	size := int((radius + thickness) * 2)
	img := ebiten.NewImage(size, size)

	vector.StrokeCircle(img, radius+thickness, radius+thickness, radius, thickness, clr, true)
	return img
}

func DrawCircleText(content string, width, height float32) *baseEntity.Drawable {
	return baseEntity.NewDrawable(ebiten.NewImage(int(width), int(height))).
		DrawImage(
			DrawText(content, basicfont.Face7x13),
			baseEntity.NewDrawOptions().
				SetScale(4, 4).
				SetPosition(float64(width/2), float64(height/2)),
		).
		DrawImage(
			DrawCircleOutline(40, 4),
			baseEntity.NewDrawOptions().
				SetPosition(float64(width/2), float64(height/2)),
		)
}
