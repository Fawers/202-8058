package main

import (
	"fmt"
	"os"
	"strconv"
)

func getDiaDaSemana(value int) string {
	// switch-case que retorna o nome do dia da semana
	// value == 0 => Domingo
	// value == 1 => Segunda
	//... value == 6 => Sábado
	switch value {
	case 0:
		return ("Domingo")
	case 1:
		return ("Segunda")
	case 2:
		return ("Terça")
	case 3:
		return ("Quarta")
	case 4:
		return ("Quinta")
	case 5:
		return ("Sexta")
	case 6:
		return ("Sábado")
	}

	return ""
}

func main() {
	value, _ := strconv.Atoi(os.Args[1])

	fmt.Println("Dia da semana: ", getDiaDaSemana(value))
}
