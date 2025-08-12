package main

import (
	"fmt"
)

func funcao(n int) string {
	if n%2 == 0 {
		return "par"
	}
	return "ímpar"
}

func main() {
	fmt.Println("Criar uma função que receba um número e retorne se ele é par ou ímpar.")

	fmt.Println(4, funcao(4))
	fmt.Println(7, funcao(7))
}
