package main

import "fmt"

// type Predicate interface {
// 	Check(bool) bool
// }
type Predicate = func(bool) bool

func newBoolIterator(iterations int) <-chan *bool {
	c := make(chan *bool)
	t, f := true, false

	go func() {
		for i := 0; i < iterations; i++ {
			switch i % 3 {
			case 0:
				c <- nil

			case 1:
				c <- &f

			case 2:
				c <- &t
			}
		}
		close(c)
	}()

	return c
}

func filterNil(iter <-chan *bool) <-chan bool {
	c := make(chan bool)

	go func() {
		for v := range iter {
			if v != nil {
				c <- *v
			}
		}
		close(c)
	}()

	return c
}

func filter(iter <-chan bool, predicate Predicate) <-chan bool {
	c := make(chan bool)

	go func() {
		for v := range iter {
			if predicate(v) {
				c <- v
			}
		}
		close(c)
	}()

	return c
}

// type OnlyFalse struct{}

// func (OnlyFalse) Check(b bool) bool {
// 	return !b
// }

func onlyFalse(b bool) bool {
	return !b
}

func main() {
	iter := filterNil(newBoolIterator(5))

	var p Predicate = func(b bool) bool {
		return b
	}
	iter = filter(iter, p)

	// iter = filter(iter, func(b bool) bool { return b })

	for v := range iter {
		fmt.Printf("%v\n", v)
	}
}
