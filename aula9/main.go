package main

import (
	"4cache/routes"
	"fmt"
	"net/http"
)

func main() {
	address := "localhost:8000"
	routes4cache := routes.GetRoutes()

	http.Handle("/4cache/", http.StripPrefix("/4cache", routes4cache))
	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println("não foi possível iniciar o servidor")
		fmt.Printf("erro: %s\n", err)
	}
}
