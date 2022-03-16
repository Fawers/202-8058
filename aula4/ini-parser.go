package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IniRecord = map[string]string
type IniMap = map[string]IniRecord

func parseIni(file *os.File) IniMap {
	// {
	// 	"202go": {
	// 		"turma": 8058,
	// 		"participantes": nomes
	// 	}
	// }

	ini := make(IniMap)
	header := ""
	ini[header] = make(IniRecord)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.IndexRune(line, '=') != -1 {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			ini[header][key] = val

		} else if strings.HasPrefix(line, "[") {
			i := 1
			j := strings.IndexRune(line, ']')
			header = line[i:j]
			ini[header] = make(IniRecord)
		}
	}

	file.Close()
	return ini
}

func printIni(ini IniMap) {
	for header, record := range ini {
		fmt.Printf("[%s]\n", header)

		for key, val := range record {
			fmt.Printf("%s = %s\n", key, val)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Necessário passar o nome do arquivo como argumento")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Não consegui ler o arquivo. erro: %s\n", err)
		return
	}

	ini := parseIni(file)
	printIni(ini)
}
