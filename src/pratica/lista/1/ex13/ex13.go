package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Criar uma função que receba um array de strings e retorne o maior texto."
	fmt.Println(text)

	list := strings.Split(text, " ")

	var gt string

	for i, value := range list {
		_gt := []rune(gt)
		_value := []rune(value)
		if len(_value) > len(_gt) {
			gt = value
			continue
		}

		if i-1 == len(list) {
			break
		}
	}

	fmt.Println("A maior palavra é:", gt)
}
