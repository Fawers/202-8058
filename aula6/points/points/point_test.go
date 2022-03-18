package points

import "testing"

func TestNewPoint(t *testing.T) {
	expected := Point2D{1, 2}
	actual := New(1, 2)

	if actual != expected {
		t.Errorf("actual (%#v) != expected (%#v)", actual, expected)
	}
}

func TestPointAdd(t *testing.T) {
	p1, p2 := New(2, 3), New(3, 2)
	expected := New(5, 5)
	actual := p1.Add(&p2)

	if actual != expected {
		t.Errorf("actual (%#v) != expected (%#v)", actual, expected)
	}
}

func TestPointAddOriginIsSamePoint(t *testing.T) {
	p, o := New(10, 15), NewOrigin()
	expected := New(10, 15)
	actual := p.Add(&o)

	if actual != expected {
		t.Errorf("actual (%#v) != expected (%#v)", actual, expected)
	}
}

func TestPointSub(t *testing.T) {
	p1, p2 := New(10, 10), New(3, 7)
	expected := New(7, 3)
	actual := p1.Sub(&p2)

	if actual != expected {
		t.Errorf("actual (%#v) != expected (%#v)", actual, expected)
	}
}

func TestPointMul(t *testing.T) {
	p := New(15, 0)
	expected := New(30, 0)
	actual := p.Mul(2)

	if actual != expected {
		t.Errorf("actual (%#v) != expected (%#v)", actual, expected)
	}
}
