package main

import "fmt"

func soma(nums [10]int) int {
	var r int
	// r := 0

	for _, num := range nums {
		r += num
	}

	return r
}

func main() {
	var nums [10]int
	nums[5] = 10

	fmt.Printf("Soma dos números: %d\n", soma(nums))
}

func main3() {
	// var nomes [5]string = [...]string{
	// 	"Fabricio", "Zé", "Elimar", "Kaike", "Marcelo",
	// }
	nomes := [...]string{
		"Marcelo",
		"Kaike",
		"Fabricio",
		"Elimar",
		"Zé",
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}
}

func main2() {
	var nomes [5]string

	nomes[0] = "Fabrício"
	nomes[1] = "Elimar"
	nomes[2] = "Zé"
	nomes[3] = "Marcelo"
	nomes[4] = "Kaike"

	for i := 0; i < len(nomes); i++ {
		fmt.Printf("nomes[%d] => %s\n", i, nomes[i])
	}

	for i, nome := range nomes {
		fmt.Printf("nomes[%d] => %s\n", i, nome)
	}
}
