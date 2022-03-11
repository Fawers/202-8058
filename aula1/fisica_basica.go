package main

import "fmt"

func calc_avg_speed(dist, time float64) float64 {
	return dist / time
}

func main() {
	dist := 130.0
	time := 1.5

	fmt.Printf("dist: %.2f, time: %.2f --> avg speed: %.2f\n",
		dist, time, calc_avg_speed(dist, time))
}
