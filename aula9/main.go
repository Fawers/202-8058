package main

import (
	"4cache/routes"
	"fmt"
	"net/http"
)

func main() {
	address := "localhost:8000"

	r := routes.NewRoutes()
	r.Cd().Add("curso", "go")
	r.Cd().Add("turma", "8058")
	r.Cd().Add("mês", "março")

	routes4cache := routes.GetRoutes(r)

	http.Handle("/4cache/", http.StripPrefix("/4cache", routes4cache))
	fmt.Printf("Inicializando servidor em %s...\n", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println("não foi possível iniciar o servidor")
		fmt.Printf("erro: %s\n", err)
	}
}
