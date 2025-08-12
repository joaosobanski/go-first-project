package main

import (
	"fmt"
)

func main() {
	fmt.Println("Criar um array de 5 números inteiros, atribuir valores manualmente e imprimir todos eles.")

	list := []int{1, 2, 3, 4, 5}

	for {
		var newNumber int
		fmt.Print("Digite um número: ")
		_, err := fmt.Scan(&newNumber)
		if err != nil {
			fmt.Printf(err.Error())
			fmt.Println("Entrada inválida, tente novamente.")
			continue
		}

		fmt.Println("Você digitou:", newNumber)
		list = append(list, newNumber)

		fmt.Println(list)
		break
	}
}
