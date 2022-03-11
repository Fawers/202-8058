package main

import "fmt"

func loopAtoB(c chan<- int, a, b int) {
	for i := a; i < b; i++ {
		fmt.Println("Enviei: ", i)
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)

	go loopAtoB(c, 0, 10)

	for v := range c {
		fmt.Println("Recebi: ", v)
	}
}
