package main

import (
	"fmt"
)

func calcMedia(list []int) float64 {
	var total int
	for _, number := range list {
		total += number
	}
	fmt.Println(total, len(list))

	return float64(total) / float64(len(list))
}

func main() {
	fmt.Println("Criar uma função que receba um array de inteiros e retorne a média.")
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(calcMedia(list))
}
