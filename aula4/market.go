package main

type Fruit string
type Price float64
type Market map[Fruit]Price
type Basket map[Fruit]uint

func buildMarket() Market {
	return nil
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
}
