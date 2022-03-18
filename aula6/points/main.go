package main

import (
	"fmt"
	"points/points"
)

func main() {
	origin := points.NewOrigin()
	point := points.New(2, 3)
	point2 := points.New(5, 10)

	addedPoint := origin.Add(&point)
	subbedPoint := point.Sub(&point2)

	fmt.Printf("%v\n", point == addedPoint)
	fmt.Printf("%v\n", subbedPoint == points.New(-3, -7))
}
