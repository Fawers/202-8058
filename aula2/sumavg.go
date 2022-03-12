package main

import "fmt"

func calcSumAvg(nums ...int) (int, float64) {
	if len(nums) == 0 {
		return 0, 0
	}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	avg := float64(sum) / float64(len(nums))

	return sum, avg
}

func main() {
	s, a := calcSumAvg(3, 4, 5)
	fmt.Printf("Soma: %d, Média: %.2f\n", s, a)

	s, a = calcSumAvg([]int{3, 4, 5, 6, 7, 8, 9, 1, 2, -5, -7}...)
	fmt.Printf("Soma: %d, Média: %.2f\n", s, a)

	s, a = calcSumAvg()
	fmt.Printf("Soma: %d, Média: %.2f\n", s, a)
}
