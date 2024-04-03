package main

import (
	"fmt"

	"github.com/paulormbsd/zeus-hammer/models"
	"github.com/paulormbsd/zeus-hammer/routes"
)

// Função main para testar, cria uma mensagem de iniciar o serviço Zeus Hammer e chama a função de rota no final
func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando Zeus Hammer")
	routes.HandleRequest()
}
