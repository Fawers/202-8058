package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "um pedaço de texto entre aspas duplas"
	var r rune = 'r'
	var n rune = 97

	fmt.Printf("%clô, %conaldo\n", n, r)
	fmt.Println(strings.Compare(s, "zum pedaço"))
	println(strings.Contains(s, "pedaço"))
	println(strings.Contains(s, "pedaco"))
	println(strings.ContainsAny(s, "psa"))
	println(strings.ContainsRune(s, '>'))
	fmt.Printf("%#v\n", strings.Fields(s))
	fmt.Printf(
		"%#v\n",
		strings.FieldsFunc(s, func(r rune) bool { return r == 'p' }))
	fmt.Printf("%#v\n", strings.Join([]string{"olá", "mundo"}, " "))
}
