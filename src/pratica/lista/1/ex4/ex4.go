package main

import (
	"fmt"
)

func main() {
	fmt.Println("Encontrar o maior e o menor número em um array de inteiros.")

	list := []int{5, 123, 3, 5, 13, 321, 3, 2, 1, 3, 2}

	var gt int

	for _, number := range list {
		if number > gt {
			gt = number
			continue
		}
	}
	fmt.Println("Maior número do slice é:", gt)
}
