package main

import (
	"fmt"

	"github.com/LeonardoBatistaCarias/go-api-rest/database"
	"github.com/LeonardoBatistaCarias/go-api-rest/models"
	"github.com/LeonardoBatistaCarias/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Leonardo", Historia: "Historia Leonardo"},
		{Id: 2, Nome: "Batista", Historia: "Historia Batista"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
