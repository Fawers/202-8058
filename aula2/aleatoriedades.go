package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func getRandNumber(lo, hi int) int {
	rand.Seed(time.Now().UnixMicro())
	num := rand.Int()

	if lo <= num && num <= hi {
		return num
	}

	validValues := hi - lo
	num %= (validValues + 1)
	return num + lo
}

func getLoHiFromArgs() (int, int) {
	if len(os.Args) < 3 {
		return -1, -1
	}

	lo, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return -1, -1
	}

	hi, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return -1, -1
	}

	return lo, hi
}

func guessingGame(lo, hi, value, tries int) bool {
	var guess int
	victory := false

	for counter := 1; counter <= tries; {
		fmt.Printf("Digite um número entre %d e %d\n", lo, hi)
		_, err := fmt.Scanf("%d", &guess)

		if err != nil {
			fmt.Printf("Não entendi - %s\n", err)
			continue
		}

		if guess < value {
			fmt.Println("Número muito baixo")
		} else if guess > value {
			fmt.Println("Número muito alto")
		} else {
			victory = true
			break
		}

		fmt.Printf("Tentativas restantes: %d\n\n", tries-counter)
		counter++
	}

	return victory
}

func main() {
	a, b := getLoHiFromArgs()

	if b <= a {
		fmt.Println("Lo e Hi são inválidos. Necessário: Hi > Lo")
		return
	}

	n := getRandNumber(a, b)

	if guessingGame(a, b, n, 5) {
		fmt.Println("Você acertou!")
	} else {
		fmt.Printf("Você perdeu! O número era %d. :(\n", n)
	}
}
