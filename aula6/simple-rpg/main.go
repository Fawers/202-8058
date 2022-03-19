package main

import (
	"fmt"
	"rpg/dice"
)

func main() {
	d := dice.NewLoaded(6)

	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
	fmt.Println(d.Roll())
}
