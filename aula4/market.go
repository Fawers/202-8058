package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Fruit = string
type Price = float64
type Market map[Fruit]Price
type Basket map[Fruit]uint

func buildMarket(filename string) (Market, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	content := string(bytes)
	market := make(Market)

	for _, line := range strings.Split(content, "\n") {
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			fmt.Printf("linha não possui sinal de igual: %q\n", line)

		} else {
			fruit := strings.TrimSpace(parts[0])
			priceStr := strings.TrimSpace(parts[1])

			price, err := strconv.ParseFloat(priceStr, 64)
			if err != nil {
				fmt.Printf("não consegui converter para float: %q\n", priceStr)
				continue
			}

			market[fruit] = price
		}
	}

	return market, nil
}

func mainMenu(market Market) {

}

func main() {
	// receber nome do arquivo por linha de comando
	// ler o arquivo e montar a quitanda, lidando com erros
	// ler input do usuário para conseguir
	// 1. ver os itens da quitanda - frutas e seus preços
	// 2. adicionar itens à cesta (Basket)
	// 2.1. escolher fruta para adicionar
	// 2.2. opcionalmente inserir uma quantidade
	// 3. calcular o preço total
	// 4. finalizar o programa quando usuário digitar q
	if len(os.Args) < 2 {
		fmt.Println("por favor passe um nome de arquivo na linha de comando")
		return
	}

	filename := os.Args[1]
	market, err := buildMarket(filename)

	if err != nil {
		fmt.Printf("erro ao construir a quitanda: %s\n", err)
		return
	} else {
		fmt.Printf("%#v\n", market)
	}

	mainMenu(market)
}
