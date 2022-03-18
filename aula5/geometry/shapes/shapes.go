package shapes

import "geometry/shapes/types"

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

func NewRectangle(length, width float64) *types.Rectangle {
	return &types.Rectangle{Length: length, Width: width}
}

func NewSquare(size float64) *types.Rectangle {
	return NewRectangle(size, size)
}

func NewTriangle(base, left, right, height float64) *types.Triangle {
	return &types.Triangle{
		Base:   base,
		Left:   left,
		Right:  right,
		Height: height,
	}
}

func NewEllipse(semimajor, semiminor float64) *types.Ellipse {
	return &types.Ellipse{
		Semimajor: semimajor,
		Semiminor: semiminor,
	}
}

func NewCircle(radius float64) *types.Ellipse {
	return NewEllipse(radius, radius)
}

func init() {
	var rec Shape2D = &types.Rectangle{}
	var tri Shape2D = &types.Triangle{}
	var cir Shape2D = &types.Ellipse{}

	rec.Area()
	tri.Area()
	cir.Area()
}
