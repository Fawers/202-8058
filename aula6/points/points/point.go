package points

type Point2D struct {
	X, Y int
}

func New(x, y int) Point2D {
	return Point2D{X: x, Y: y}
}

func NewOrigin() Point2D {
	return New(0, 0)
}

func (p *Point2D) Add(p2 *Point2D) (r Point2D) {
	r.X = p.X + p2.X
	r.Y = p.Y + p2.Y
	return
}

func (p *Point2D) Sub(p2 *Point2D) (r Point2D) {
	r.X = p.X - p2.X
	r.Y = p.Y - p2.Y
	return
}

func (p *Point2D) Mul(multiplier int) (r Point2D) {
	r.X = p.X * multiplier
	r.Y = p.Y * multiplier
	return
}
