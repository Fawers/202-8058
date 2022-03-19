package dice

import (
	"math/rand"
	"time"
)

type Die interface {
	Roll() uint
	GetMax() uint
}

type RegularDie struct {
	max uint
	r   *rand.Rand
}

func New(max uint) Die {
	d := new(RegularDie)
	d.max = max
	d.r = rand.New(rand.NewSource(time.Now().UnixMicro()))
	return d
}

func (d *RegularDie) Roll() uint {
	n := uint(d.r.Uint32())
	return n % d.max
}

func (d *RegularDie) GetMax() uint {
	return d.max - 1
}

func NewLoaded(value uint) Die {
	d := new(LoadedDie)
	d.value = value
	return d
}

type LoadedDie struct {
	value uint
}

func (d *LoadedDie) Roll() uint {
	return d.value
}

func (d *LoadedDie) GetMax() uint {
	return d.value
}
