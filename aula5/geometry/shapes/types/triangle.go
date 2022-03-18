package types

type Triangle struct {
	Base, Left, Right float64
	// Height refers to the perpendicular height
	Height float64
}

func (t *Triangle) IsEquilateral() bool {
	return t.Base == t.Left && t.Base == t.Right
}

func (t *Triangle) IsIsosceles() bool {
	return (!t.IsEquilateral() &&
		t.Base == t.Left ||
		t.Base == t.Right ||
		t.Left == t.Right)
}

func (t *Triangle) IsScalene() bool {
	return !t.IsEquilateral() && !t.IsIsosceles()
}

func (t *Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func (t *Triangle) Perimeter() float64 {
	return t.Base + t.Left + t.Right
}
