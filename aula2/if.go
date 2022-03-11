package main

import (
	"fmt"
	"os"
	"strconv"
)

func is_age_major(age uint8) bool {
	return age >= 18
}

func outro_main() {
	if len(os.Args) < 2 {
		fmt.Printf("Uso:\n\t%s IDADE:uint8\n", os.Args[0])
		return
	}

	num, err := strconv.ParseUint(os.Args[1], 10, 8)

	if err != nil {
		fmt.Printf("erro: %s\n", err)
		return
	}

	if is_age_major(uint8(num)) {
		fmt.Println("É maior de idade -> bora pro bar")
	} else {
		fmt.Println("Não é maior de idade -> pega um toddyinho")
	}
}

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("maior=")
	}
}
