package model

type Triangle struct {
	Base      float32
	Height    float32
	SideLeft  float32
	SideRight float32
}

func (t *Triangle) Area() float32 {
	return (t.Base * t.Height) / 2
}

func (t *Triangle) Perimeter() float32 {
	return t.SideLeft + t.SideRight + t.Base
}
