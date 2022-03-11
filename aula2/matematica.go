package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Valor de PI: ", math.Pi)
	fmt.Println("Valor de E: ", math.E)

	for i := 1; i <= 10; i++ {
		i := float64(i)
		h := i / 2

		fmt.Printf("Y0(%.1f) = %f\n", h, math.Y0(h))
		fmt.Printf("Y0(%.1f) = %f\n", i, math.Y0(i))
	}

	fmt.Println(math.Abs(-10))
	fmt.Println("raiz de 25: ", math.Sqrt(25))
}
