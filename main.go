package main

import (
	"github.com/WeslleySanto/go-validacoes-testes-paginas-html/database"
	"github.com/WeslleySanto/go-validacoes-testes-paginas-html/routes"
)

func main() {
	database.Conecta()
	routes.HandleRequests()
}
