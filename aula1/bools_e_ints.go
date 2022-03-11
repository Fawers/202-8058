package main

import "fmt"

func valores_booleanos() {
	var verdadeiro bool = true
	falso := false
	fmt.Println(verdadeiro, falso)
}

func valores_inteiros() {
	//          size_t
	var inteiro int = 0x7fffffffffffffff
	var i8bits int8 = -128
	var i16bits int16 = 0x7fff
	var i32bits int32 = 345678765
	var i64bits int64 = -9_000_000_000_000_000_000
	fmt.Println(inteiro)
	fmt.Printf("8bits:  %d\n16bits: %d\n32bits: %d\n64bits: %d\n",
		i8bits, i16bits, i32bits, i64bits)

	//                    usize_t
	var inteiro_sem_sinal uint = 0xffffffffffffffff
	var u8 uint8 = 255
	var u32 uint32 = 0xffffffff
	fmt.Println(inteiro_sem_sinal)
	fmt.Printf("8bits:  %d\n32bits: %d\n", u8, u32)
}

func main() {
	valores_booleanos()
	valores_inteiros()
}
