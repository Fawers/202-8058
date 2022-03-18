package types

type Rectangle struct {
	Length, Width float64
}

func (r *Rectangle) IsSquare() bool {
	return r.Length == r.Width
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.Length + 2*r.Width
}
