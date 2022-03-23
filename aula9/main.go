package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá, mundão!")
}

func mkHelloX(queryParam string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := "world"
		query := r.URL.Query()
		if query.Has(queryParam) {
			name = query.Get(queryParam)
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

type helloXContext struct {
	queryParam string
}

func (ctx *helloXContext) handle(w http.ResponseWriter, r *http.Request) {
	name := "world"
	query := r.URL.Query()
	if query.Has(ctx.queryParam) {
		name = query.Get(ctx.queryParam)
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func helloX(w http.ResponseWriter, r *http.Request) {
	name := "world"
	query := r.URL.Query()
	if query.Has("name") {
		name = query.Get("name")
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func cacheHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home do 4Cache")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/hello", helloX)
	http.HandleFunc("/hello/nome", mkHelloX("nome"))
	http.HandleFunc("/hello/nimi", (&helloXContext{"nimi"}).handle)

	sm4cache := http.NewServeMux()
	sm4cache.HandleFunc("/", cacheHome)

	http.Handle("/4cache", sm4cache)

	address := "localhost:8000"
	fmt.Printf("Inicializando servidor em %s...\n", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println("Não foi possível iniciar o servidor")
		fmt.Printf("erro: %s\n", err)
	}
}
