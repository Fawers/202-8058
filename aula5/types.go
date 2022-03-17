package main

import (
	"fmt"
	"os"
)

type Dog struct {
	Name  string
	Breed string // raça
	Age   uint8
}

func (d *Dog) print() {
	fmt.Printf(
		"%s é um %s e tem %d anos e se encontra em %p.\n",
		d.Name, d.Breed, d.Age, d)
}

func (d *Dog) Bark() {
	fmt.Printf("%s está latindo! Au au!\n", d.Name)
}

func (d *Dog) MakeBirthday() {
	d.Age++ // aumenta o número em 1
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("passe o nome e a raça do cachorro via linha de comando")
		return
	}

	dogName := os.Args[1]
	dogBreed := os.Args[2]
	var dog Dog = Dog{dogName, dogBreed, 0}

	fmt.Printf("%#v :: %p\n", dog, &dog)

	dog.print()
	dog.Bark()

	dog.MakeBirthday()
	dog.print()
	dog.Bark()

	fmt.Printf("%#v :: %p\n", dog, &dog)
}
