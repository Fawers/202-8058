package main

import (
	"fmt"
	"os"
)

// import fmt
// fmt.go -> import os
// os.go -> import xpto
// xpto.go -> 0 imports
// comp xpto.go
// comp os.go <-> xpto.o
// comp fmt.go <-> os.o
// comp main.go <-> fmt.o

func mkIntList(create bool) []int {
	var slice []int

	if create {
		slice = make([]int, 0)
		// outras coisas com esse slice
	}

	return slice
}

func readFile() string {
	f, err := os.Open("curso.txt")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	buffer := make([]byte, 1024)
	_, err = f.Read(buffer)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	content := string(buffer)

	return content
}

func add(a, b int) int {
	return a + b
}

func xpto() *os.File {
	return new(os.File)
}

func main() {
	fmt.Println("Olá, mundo!")
	fmt.Printf("2 + 3 = %d\n", add(2, 3))
	fmt.Print("Conteúdo do arquivo:", readFile(), "\n")
	//file := xpto()
	// <--
}
