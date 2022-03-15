package main

import "fmt"

// fatorial(0 | 1) = 1
// fatorial(x), x > 1 = x * fatorial(x-1)

func factorial_switch(num uint64) uint64 {
	switch num {
	case 0, 1:
		return 1
	case 2:
		return 2
	case 3:
		return 6
	default:
		return num * factorial_switch(num-1)
	}
}

func factorial_ifelse(num uint64) uint64 {
	if num == 0 || num == 1 {
		return 1
	}

	return num * factorial_ifelse(num-1)
	// 5 * factorial(4)
	// 4 * factorial(3)
	// 3 * factorial(2)
	// 2 * factorial(1)
	// 2 * 1
	// 3 * 2 = 6
	// 4 * 6 = 24
	// 5 * 24 = 120
}

func main() {
	var i uint64
	for i = 0; i <= 3; i++ {
		fmt.Printf("%2d! = %7d\n", i, factorial_ifelse(i))
	}

	for i = 4; i < 7; i++ {
		fmt.Printf("%2d! = %7d\n", i, factorial_switch(i))
	}
}
