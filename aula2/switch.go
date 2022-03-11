package main

import (
	"fmt"
	"os"
	"strconv"
)

func getIntArg(index int) int {
	if len(os.Args) <= index {
		return -1
	}

	n, err := strconv.Atoi(os.Args[index])
	if err != nil {
		return -1
	}

	return n
}

func main() {
	m := getIntArg(1)
	n := getIntArg(2)

	switch m % n {
	case 0:
		fmt.Printf("%d é divisível por %d\n", m, n)

	case 1:
		fmt.Printf("%d não é divisível por %d por 1 unidade\n", m, n)

	case 2:
		fmt.Printf("%d não é divisível por %d por 2 unidades\n", m, n)
		fallthrough

	default:
		fmt.Println("yabadabadoo")
	}
}
