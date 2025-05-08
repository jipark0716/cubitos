package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable struct {
	Image  *ebiten.Image
	Option ebiten.DrawImageOptions
}

func NewDrawable(image *ebiten.Image) *Drawable {
	return &Drawable{
		Image:  image,
		Option: ebiten.DrawImageOptions{},
	}
}

func (d *Drawable) Copy() *Drawable {
	return &Drawable{
		Image:  d.Image,
		Option: d.Option,
	}
}

func (d *Drawable) CopyWithClear() *Drawable {
	res := d.Copy()
	res.Image.Clear()
	return res
}

func (d *Drawable) SetCenterAnchor() {
	size := d.Image.Bounds().Size()
	w, h := float64(size.X/2), float64(size.Y/2)
	d.Option.GeoM.Translate(-w, -h)
}

func (d *Drawable) SetStartAnchor() {
	size := d.Image.Bounds().Size()
	w, h := float64(size.X/2), float64(size.Y/2)
	d.Option.GeoM.Translate(w, h)
}

func (d *Drawable) Draw(draw *Drawable) {
	d.Image.DrawImage(draw.Image, &draw.Option)
}

func (d *Drawable) DrawImage(img *ebiten.Image, opt *DrawOptions) *Drawable {
	dest := NewDrawable(img)
	dest.SetCenterAnchor()
	opt.Translate(&dest.Option)
	d.Image.DrawImage(dest.Image, &dest.Option)

	return d
}

func (d *Drawable) Translate(opt *DrawOptions) *Drawable {
	opt.Translate(&d.Option)

	return d
}
