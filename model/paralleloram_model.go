package model

type Parallelogram struct {
	Base   float32
	Height float32
	Side   float32
}

func (p *Parallelogram) Area() float32 {
	return p.Base * p.Height
}

func (p *Parallelogram) Perimeter() float32 {
	return 2 * (p.Base + p.Side)
}
