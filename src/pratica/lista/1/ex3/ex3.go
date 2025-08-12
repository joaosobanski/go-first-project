package main

import (
	"fmt"
)

func main() {
	fmt.Println("Somar todos os elementos de um array de inteiros.")

	list := []int{1, 2, 3, 4, 5}

	total := 0

	for i, number := range list {
		fmt.Println(i, number)
		total += number

	}

	fmt.Println("Total:", total)
}
