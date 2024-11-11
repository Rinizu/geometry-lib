package model

type Rectangle struct {
	Width  float32
	Length float32
}

func (r *Rectangle) Area() float32 {
	return r.Width * r.Length
}

func (r *Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Width)
}
