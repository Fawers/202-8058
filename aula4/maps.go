package main

import "fmt"

func main() {
	// {chave: "valor", chave2: "valor2"}

	letras := make(map[rune]int)
	nomes := "Fabricio,Elimar,Kaike,Zé,Marcelo,Thiago"

	for _, c := range nomes {
		letras[c] += 1
	}

	for key, val := range letras {
		fmt.Printf("letras[%c]:%d\n", key, val)
	}
}
