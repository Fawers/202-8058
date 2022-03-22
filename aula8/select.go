package main

import "fmt"

func fibonacci(stop <-chan bool) <-chan uint64 {
	var a, b uint64 = 0, 1
	c := make(chan uint64)

	go func() {
		running := true
		for running {
			select {
			case c <- a:
				a, b = b, a+b

			case stopRunning := <-stop:
				running = !stopRunning
			}
		}

		close(c)
	}()

	return c
}

func main() {
	stopper := make(chan bool)
	fibIterator := fibonacci(stopper)

	for fib := range fibIterator {
		fmt.Printf("recebi o valor %d\n", fib)

		stopper <- fib >= 100000
	}
}
