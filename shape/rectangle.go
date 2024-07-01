package shape

type Rectangle struct {
	Width  float32
	Length float32
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.Width * rectangle.Length
}

func (rectangle *Rectangle) Perimeter() float32 {
	return 2 * (rectangle.Length + rectangle.Width)
}
