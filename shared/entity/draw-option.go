package entity

import "github.com/hajimehoshi/ebiten/v2"

type DrawOptions struct {
	scaleX, scaleY         *float64
	posX, posY             *float64
	clrR, clrG, clrB, clrA *float32
}

func NewDrawOptions() *DrawOptions {
	return &DrawOptions{}
}

func (d *DrawOptions) Translate(opt *ebiten.DrawImageOptions) {
	if d.scaleX != nil && d.scaleY != nil {
		opt.GeoM.Scale(*d.scaleX, *d.scaleY)
	}

	if d.posX != nil && d.posY != nil {
		opt.GeoM.Translate(*d.posX, *d.posY)
	}

	if d.clrR != nil && d.clrG != nil && d.clrB != nil && d.clrA != nil {
		opt.ColorScale.Scale(
			*d.clrR/255,
			*d.clrG/255,
			*d.clrB/255,
			*d.clrA/255,
		)
	}
}

func (d *DrawOptions) SetScale(x, y float64) *DrawOptions {
	d.scaleX = &x
	d.scaleY = &y
	return d
}

func (d *DrawOptions) SetPosition(x, y float64) *DrawOptions {
	d.posX = &x
	d.posY = &y
	return d
}

func (d *DrawOptions) SetColorMask(r, g, b, a float32) *DrawOptions {
	d.clrR = &r
	d.clrG = &g
	d.clrB = &b
	d.clrA = &a
	return d
}
