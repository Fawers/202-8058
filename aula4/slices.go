package main

import "fmt"

func sum(nums []int) (result int) {
	fmt.Printf("nums in sum: %v\n", nums)

	for _, val := range nums {
		result += val
	}

	return
}

func main() {
	nums := [...]int{
		1, 2, 3, 4, 5,
	}
	numsSlice := nums[1 : len(nums)-1]

	fmt.Printf("Soma de %v: %d\n", numsSlice, sum(numsSlice))

	numsSlice[1] = 7
	fmt.Printf("arr: %v\nslc: %v\n", nums, numsSlice)

	// numsSlice[3] = 10 ==> erro: OutOfBounds

	fmt.Printf("Endereço do array: %p\n", &nums[1])
	fmt.Printf("Endereço do slice: %p\n", numsSlice)

	numsSlice = append(numsSlice, 10)
	fmt.Printf("arr: %v\nslc: %v\n", nums, numsSlice)
}

func main2() {
	nomes := []string{
		"Elimar",
		"Kaike",
		"Zé",
		"Kaike",
		"Marcelo",
	}

	fmt.Printf("endereço   de nomes: %p\n", nomes)
	fmt.Printf("len        de nomes: %d\n", len(nomes))
	fmt.Printf("capacidade de nomes: %d\n", cap(nomes))

	nomes = append(nomes, "Fabrício")

	fmt.Printf("\nendereço   de nomes: %p\n", nomes)
	fmt.Printf("len        de nomes: %d\n", len(nomes))
	fmt.Printf("capacidade de nomes: %d\n", cap(nomes))

	nomes = append(nomes, "Thiago")

	fmt.Printf("\nendereço   de nomes: %p\n", nomes)
	fmt.Printf("len        de nomes: %d\n", len(nomes))
	fmt.Printf("capacidade de nomes: %d\n", cap(nomes))
}
