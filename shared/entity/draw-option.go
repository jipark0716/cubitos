package entity

type DrawOptions struct {
	scaleX, scaleY float64
	posX, posY     float64
}

func NewDrawOptions() *DrawOptions {
	return &DrawOptions{
		scaleX: 1,
		scaleY: 1,
		posX:   0,
		posY:   0,
	}
}

func (d *DrawOptions) SetScale(x, y float64) *DrawOptions {
	d.scaleX = x
	d.scaleY = y
	return d
}

func (d *DrawOptions) SetPosition(x, y float64) *DrawOptions {
	d.posX = x
	d.posY = y
	return d
}
