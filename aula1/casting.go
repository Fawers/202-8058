package main

import "fmt"

func main() {
	var idade uint8 = 29
	idade_quebrada := float32(idade)

	fmt.Printf("%d %f\n", idade, idade_quebrada)

	s := string(0x41)
	fmt.Println("<", s, ">")
	// i := uint32(s)
}
