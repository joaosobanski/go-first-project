package main

import (
	"fmt"
)

func calculate(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	fmt.Println("Criar uma função que receba dois números e retorne dois valores: a soma e a multiplicação.")

	soma, multiplicacao := calculate(3, 4)

	fmt.Println("Soma:", soma)
	fmt.Println("Multiplicação:", multiplicacao)
}
