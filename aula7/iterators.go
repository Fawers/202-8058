package main

import "fmt"

func newIntIterator(start, stop int) <-chan int {
	c := make(chan int)

	go func() {
		for i := start; i < stop; i++ {
			fmt.Printf("<- newIntIterator:%v\n", i)
			c <- i
		}
		close(c)
	}()

	return c
}

func filterOdd(iter <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for v := range iter {
			fmt.Printf("<- filterOdd:%v\n", v)
			if v%2 == 1 {
				c <- v
			}
		}

		close(c)
	}()

	return c
}

func mapSquare(iter <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for v := range iter {
			fmt.Printf("<- mapSquare:%v\n", v)
			c <- v * v
		}

		close(c)
	}()

	return c
}

func main() {
	iter := newIntIterator(0, 10)
	iter = filterOdd(iter)
	iter = mapSquare(iter)

	v := <-iter

	fmt.Println(v)

	fmt.Println("fim do programa")
}
