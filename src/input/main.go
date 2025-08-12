package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 - Iniciar checker")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair")

	var comando int
	fmt.Scan(&comando)

	switch comando {
	case 1:
		fmt.Println("Iniciando checker...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 3:
		fmt.Println("Saindo...")
	case 0:
		fmt.Println("Valor inválido. Por favor, insira um número entre 1 e 3.")
	default:
		fmt.Println("Comando inválido.")
	}
}
