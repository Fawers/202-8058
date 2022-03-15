package main

import (
	"errors"
	"fmt"
)

var ErrEmptySlice = errors.New("Slice vazio")

func sum(nums []int) (int, error) {
	switch len(nums) {
	case 0:
		return 0, ErrEmptySlice

	case 1:
		return nums[0], nil

	default:
		s, _ := sum(nums[1:])
		return nums[0] + s, nil
	}
}

func main() {
	nums := []int{
		5, 567, 346, 23, 7, 23, 67, 23, 67,
	}

	manyNums := [][]int{
		nums,
		[]int{},
	}

	for _, slice := range manyNums {
		s, err := sum(slice)

		switch err {
		case nil:
			fmt.Printf("soma de %v = %d\n", nums, s)

		case ErrEmptySlice:
			fmt.Printf("deu ruim: %s\n", err)
		}
	}
}
