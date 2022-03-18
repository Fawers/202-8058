package types

import "math"

type Ellipse struct {
	Semimajor, Semiminor float64 // axes
}

func (e *Ellipse) IsCircle() bool {
	return e.Semimajor == e.Semiminor
}

func (e *Ellipse) Area() float64 {
	return math.Pi * e.Semimajor * e.Semiminor
}

func (e *Ellipse) Perimeter() float64 {
	if e.IsCircle() {
		return 2 * math.Pi * e.Semimajor
	} else {
		a, b := e.Semimajor, e.Semiminor
		return math.Pi * (3*(a+b) - math.Sqrt((a+3*b)*(b+3*a)))
	}
}
