package main

import "fmt"

type SuperStruct struct {
	x, y int
	s    string
	b    bool
	c    complex64
	r    *bool
}

func producer2_0(loopIterations int) <-chan *SuperStruct {
	c := make(chan *SuperStruct)

	go func() {
		for i := 0; i < loopIterations; i++ {
			fmt.Printf("produzindo: %d\n", i)
			if i%2 == 0 {
				c <- &SuperStruct{}
			} else {
				c <- nil
			}
		}

		close(c)
	}()

	return c
}

func producer(c chan<- int, loopIterations int) {
	for i := 0; i < loopIterations; i++ {
		fmt.Printf("produzindo: %d\n", i)
		c <- i
	}
	close(c)
	fmt.Println("canal fechado")
}

func consumer(c <-chan *SuperStruct) {
	for i := range c {
		fmt.Printf("consumindo: %v\n", i)
	}
}

func main() {
	c := producer2_0(10)
	consumer(c)
}
