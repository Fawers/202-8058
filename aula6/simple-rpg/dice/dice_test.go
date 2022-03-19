package dice

import (
	"testing"
)

func TestNewCreatesCommonDieWithGivenMax(t *testing.T) {
	var max uint = 21
	d := New(max)

	if cd, ok := d.(*RegularDie); ok {
		if cd.max != max {
			t.Errorf("dado não está com o máximo esperado (%d): %d", max, cd.max)
		}
	} else {
		t.Errorf("New não retornou uma instância de *CommonDie: %T", d)
	}
}

func TestNewConstantCreatesConstantDieWithGivenValue(t *testing.T) {
	var v uint = 6
	d := NewLoaded(v)

	if cd, ok := d.(*LoadedDie); ok {
		if cd.value != v {
			t.Errorf("valor constante diferente do esperado (%d): %d", v, cd.value)
		}
	} else {
		t.Errorf("NewConstant não retornou uma instância de *ConstantDie: %T", d)
	}
}

func TestCommonDieRollsDifferentNumbers(t *testing.T) {
	var m uint = 6
	d := New(m)

	first := d.Roll()
	diff := false

	for i := 0; i < 100; i++ {
		r := d.Roll()

		if !diff && r != first {
			diff = true
		}

		if r >= m {
			t.Fatalf("dado rolou um número maior ou igual do que max (%d): %d", m, r)
		}
	}

	if !diff {
		t.Errorf("dado não rolou nenhum número diferente de %d", first)
	}
}

func TestConstantDieAlwaysRollsSameNumber(t *testing.T) {
	var v uint = 4
	d := NewLoaded(v)

	for i := 0; i < 100; i++ {
		if r := d.Roll(); r != v {
			t.Fatalf(
				"dado viciado não rolou valor esperado de %d; rolou %d",
				v, r)
		}
	}
}
