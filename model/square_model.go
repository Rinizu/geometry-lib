package model

type Square struct {
	Side float32
}

func (s *Square) Area() float32 {
	return s.Side * s.Side
}

func (s *Square) Perimeter() float32 {
	return s.Side * 4
}
