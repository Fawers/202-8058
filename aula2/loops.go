package main

import (
	"fmt"
	"strings"
)

func is_even(n int) bool {
	return n%2 == 0
}

func string_contains_o(s string) bool {
	return strings.ContainsRune(s, 'o')
}

func main() {
	for i := 0; i < 2; i += 1 {
		if is_even(i) {
			fmt.Printf("%d é par\n", i)
		} else {
			fmt.Printf("%d é ímpar\n", i)
		}
	}

	fmt.Println()
	//     ↓↓↓↓↓↓↓↓↓↓↓↓↓ for each
	for _, name := range []string{"José", "Olaurito", "Elimar"} {
		if string_contains_o(name) {
			fmt.Printf("%s contém uma letra '%c'\n", name, 'o')
		} else {
			fmt.Printf("%s não contém uma letra '%c'\n", name, 'o')
		}
	}

	fmt.Println()
	// loop "while"
	contagem := 2
	for contagem > 0 {
		fmt.Println(contagem)
		contagem-- // contagem -= 1 <=> contagem = contagem - 1
	}

	// loop infinito
	for {
		fmt.Println("Ao infinito e além!!")
	}
}
