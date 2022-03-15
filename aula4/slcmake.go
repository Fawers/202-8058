package main

import "fmt"

func main() {
	// append_com_varios("Fabricio", "Elimar", "Kaike", "Zé", "Marcelo", "Thiago")
	forma_incomum()
}

func append_com_varios(ns ...string) {
	nomes := make([]string, 0, 2)
	fmt.Printf("len = %d, cap = %d, endereço: %p\n", len(nomes), cap(nomes), nomes)

	nomes = append(nomes, ns...)
	fmt.Printf("len = %d, cap = %d, endereço: %p\n", len(nomes), cap(nomes), nomes)
}

func forma_comum() {
	nomes := make([]string, 0, 8)
	fmt.Printf("len = %d, cap = %d, endereço: %p\n", len(nomes), cap(nomes), nomes)

	nomes = append(nomes, "Thiago")
	fmt.Printf("len = %d, cap = %d, endereço: %p\n", len(nomes), cap(nomes), nomes)
	nomes = append(nomes, "Marcelo")
	fmt.Printf("len = %d, cap = %d, endereço: %p\n", len(nomes), cap(nomes), nomes)
	nomes = append(nomes, "Elimar")
	fmt.Printf("len = %d, cap = %d, endereço: %p\n", len(nomes), cap(nomes), nomes)
	nomes = append(nomes, "Kaike")
	fmt.Printf("len = %d, cap = %d, endereço: %p\n", len(nomes), cap(nomes), nomes)
	nomes = append(nomes, "Zé")
	fmt.Printf("len = %d, cap = %d, endereço: %p\n", len(nomes), cap(nomes), nomes)
	nomes = append(nomes, "Fabrício")
	fmt.Printf("len = %d, cap = %d, endereço: %p\n", len(nomes), cap(nomes), nomes)

	for _, nome := range nomes {
		fmt.Println(nome)
	}
}

func forma_incomum() {
	nomes := make([]string, 6, 8)
	fmt.Printf("len = %d, cap = %d\n", len(nomes), cap(nomes))

	nomes[2] = "Kaike"
	nomes[4] = "Marcelo"
	nomes[5] = "Thiago"
	nomes[3] = "Zé"
	nomes[1] = "Elimar"
	nomes[0] = "Fabrício"

	for _, nome := range nomes[2:4] {
		fmt.Println(nome)
	}
}
