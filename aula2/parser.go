package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseItem(line string) (string, float64, uint, error) {
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " | ")

	if len(parts) != 3 {
		return "", 0, 0, fmt.Errorf("linha não possui 3 atributos")
	}

	name := parts[0]
	price, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return "", 0, 0, err
	}

	qty, err := strconv.ParseUint(parts[2], 10, 0)
	if err != nil {
		return "", 0, 0, err
	}

	return name, price, uint(qty), nil
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1024)
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(buf), "\n"), nil
}

func getArgs() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("Argumentos insuficientes")
	}

	return os.Args[1], nil
}

func main() {
	filename, err := getArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler nome do arquivo: %s\n", err)
		os.Exit(1)
	}

	content, err := readFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler o arquivo: %s\n", err)
		os.Exit(2)
	}

	for _, line := range content {
		name, price, qty, err := parseItem(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao parsear o item: %s\n", err)
		} else {
			fmt.Printf("Item(`%s`, %.2f, %d)\n", name, price, qty)
		}
	}
}
