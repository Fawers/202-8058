package main

import (
	"fmt"
	gtypes "go/types" // package alias (apelido de pacote)

	"geometry/shapes"
	"geometry/shapes/types"
)

func main() {
	shapes := []shapes.Shape2D{
		shapes.NewSquare(5),
		shapes.NewRectangle(4, 7),
		shapes.NewCircle(3),
		shapes.NewTriangle(4, 3, 3, 5),
		shapes.NewTriangle(5, 8, 4, 10),
		shapes.NewEllipse(8, 7),
	}

	for _, s := range shapes {
		switch s := s.(type) {
		case *types.Rectangle:
			if s.IsSquare() {
				fmt.Println("s é um quadrado!")
			} else {
				fmt.Println("s é um retângulo!")
			}

		case *types.Triangle:
			fmt.Print("s é um triângulo ")

			if s.IsEquilateral() {
				fmt.Println("equilátero!")
			} else if s.IsIsosceles() {
				fmt.Println("isósceles!")
			} else {
				fmt.Println("escaleno!")
			}

		case *types.Ellipse:
			if s.IsCircle() {
				fmt.Println("s é um círculo!")
			} else {
				fmt.Println("s é uma elípse!")
			}
		}

		fmt.Printf("área é %f\nperímetro é %f\n\n", s.Area(), s.Perimeter())
	}
}
