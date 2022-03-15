package main

import "fmt"

func main() {
	curso := "202 Go 4Linux"

	fmt.Println(curso)
	fmt.Println(curso[4:])
	fmt.Println(curso[:len(curso)-7])
	fmt.Println(curso[4 : len(curso)-7])

	for i, c := range curso {
		fmt.Printf("Caractere %d é %q\n", i, c)
	}
}
