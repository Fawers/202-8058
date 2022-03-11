package main

import (
	"fmt"
	"math/rand"
)

func getRandNumber(lo, hi int) int {
	num := rand.Int()

	// ajustar num para que lo <= num <= hi

	return num
}

func getLoHiFromArgs() (int, int) {

	return -1, -1
}

func main() {
	a, b := getLoHiFromArgs()
	n := getRandNumber(a, b)

	fmt.Println(a, b, n)
}
