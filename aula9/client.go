package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("erro ao getar %q: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("resposta: %d\n", resp.StatusCode)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("erro ao ler a resposta: %s\n", err)
		return
	}

	fmt.Print(string(content))
}
