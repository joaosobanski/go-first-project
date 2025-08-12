package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Criar uma função que receba dois números inteiros e retorne a soma deles.")
	fmt.Println(sum(1, 2))
}
