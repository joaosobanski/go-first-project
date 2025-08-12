package main

import (
	"fmt"
)

func main() {
	fmt.Println("Criar um array de strings com nomes de frutas e imprimir apenas o segundo e o quarto elemento.")

	list := []string{"apple", "banana", "pineapple", "cucumber", "orange"}

	fmt.Println("1 Fruit:", list[1])
	fmt.Println("2 Fruit:", list[3])
}
