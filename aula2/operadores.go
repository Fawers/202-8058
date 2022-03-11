package main

import (
	"fmt"
	"strings"
)

func table_and() {
	fmt.Println("p\tq\tp && q")

	for _, p := range []bool{true, false} {
		for _, q := range []bool{true, false} {
			fmt.Printf("%v\t%v\t%v\n", p, q, p && q)
			//                                 ^^ operador AND
		}
	}
}

func table_or() {
	fmt.Println("p\tq\tp || q")

	for _, p := range []bool{true, false} {
		for _, q := range []bool{true, false} {
			fmt.Printf("%v\t%v\t%v\n", p, q, p || q)
			//                                 ^^ operador OR
		}
	}
}

func ops_aritmeticos(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)

	x, y := float64(a), float64(b)
	fmt.Printf("%.2f / %.2f = %.2f\n", x, y, x/y)
}

func op_logicos(a, b int) {
	fmt.Printf("%d == %d => %v\n", a, b, a == b)
	fmt.Printf("%d != %d => %v\n", a, b, a != b)
	fmt.Printf("%d < %d => %v\n", a, b, a < b)
	fmt.Printf("%d <= %d => %v\n", a, b, a <= b)
	fmt.Printf("%d > %d => %v\n", a, b, a > b)

	fmt.Printf("%d >= %d => %v\n", a, b, a >= b)
	fmt.Printf("%d >= %d => %v\n", a, b, a > b || a == b)
}

func op_logicos_string(a, b string) {
	fmt.Printf("%s == %s => %v\n", a, b, a == b)
	fmt.Printf("%s != %s => %v\n", a, b, a != b)

	switch strings.Compare(a, b) {
	case 0:
		fmt.Printf("%s == %s\n", a, b)

	case 1:
		fmt.Printf("%s > %s\n", a, b)

	case -1:
		fmt.Printf("%s < %s\n", a, b)
	}
}

func main() {
	// table_and()
	// table_or()
	// ops_aritmeticos(6, 4)
	// op_logicos(7, 5)
	op_logicos_string("Mar", "Zé")
}
