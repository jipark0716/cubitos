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

type TitleDrawOption struct {
	perRow               int
	paddingX, paddingY   int
	gPaddingX, gPaddingY float64
}

func NewTitleDrawOption(perRow int) *TitleDrawOption {
	return &TitleDrawOption{perRow: perRow}
}

func (t *TitleDrawOption) SetPadding(padding int) *TitleDrawOption {
	t.paddingX = padding
	t.paddingY = padding
	return t
}

func (t *TitleDrawOption) SetGPadding(padding float64) *TitleDrawOption {
	t.gPaddingX = padding
	t.gPaddingY = padding
	return t
}

func (t *TitleDrawOption) SetGPaddingX(padding float64) *TitleDrawOption {
	t.gPaddingX = padding
	return t
}

func (t *TitleDrawOption) SetGPaddingY(padding float64) *TitleDrawOption {
	t.gPaddingY = padding
	return t
}

func DrawTile[T Entity](d *Drawable, targets []T, opt *TitleDrawOption) {
	for i, target := range targets {
		x := i % opt.perRow
		y := i / opt.perRow

		target.Draw(
			d.Image,
			NewDrawOptions().
				SetPosition(opt.gPaddingX+float64(x*opt.paddingX), opt.gPaddingY+float64(y*opt.paddingY)).
				SetScale(0.5, 0.5),
		)
	}
}
