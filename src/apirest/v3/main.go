package main

import (
	"fmt"
	"main/database"
	"main/model"
	"main/route"
)

func mock() {
	model.Personalidades = []model.Personalidade{
		{Nome: "Nome 1", Historia: "Historia do nome 1"},
		{Nome: "Nome 2", Historia: "Historia do nome 2"},
	}
}

func main() {
	// mock()

	database.ConectaBD()

	fmt.Println("Iniciando o servidor Rest com Go")
	route.HandleResquest()
}
