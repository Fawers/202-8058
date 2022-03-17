package main

import "fmt"

type Animal interface {
	MakeSound()
	Interact(a Animal)
	GetName() string
}

type Dog struct {
	Name string
}

func (d *Dog) MakeSound() {
	fmt.Printf("%s: auau!\n", d.Name)
}

func (d *Dog) Interact(a Animal) {
	// type assertion
	d2, ok := a.(*Dog) // downcast

	if ok {
		fmt.Printf("%s está brincando com %s!\n", d.Name, d2.Name)
	} else {
		fmt.Printf("%s está estranhando %s\n", d.Name, a.GetName())
	}
}

func (d *Dog) GetName() string {
	return d.Name
}

type Cat struct {
	Name string
}

func (c *Cat) MakeSound() {
	fmt.Printf("%s: miau!\n", c.Name)
}

func (c *Cat) Interact(a Animal) {
	// type switch
	switch a := a.(type) {
	case *Cat:
		fmt.Printf("%s está dando banho em %s\n", c.Name, a.Name)

	case *Dog:
		fmt.Printf("%s deu no pé! E está fugindo de %s\n", c.Name, a.Name)

	default:
		fmt.Printf("%s está encarando %s\n", c.Name, a.GetName())
	}
}

func (c *Cat) GetName() string {
	return c.Name
}

func main() {
	rex := Dog{"Rex"}
	thor := Dog{"Thor"}
	snowball := Cat{"Snowball"}
	popcorn := Cat{"Popcorn"}

	rex.MakeSound()
	thor.Interact(&rex)
	rex.Interact(&snowball)

	snowball.MakeSound()
	popcorn.Interact(&snowball)
	snowball.Interact(&rex)
}
