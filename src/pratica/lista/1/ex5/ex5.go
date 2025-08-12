package main

import (
	"fmt"
)

// Função para inverter letras de uma string
func inverterString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Inverter a ordem e as letras das palavras.")

	list := []string{"Inverter", "a", "ordem", "dos", "elementos", "de", "um", "array", "de", "strings"}
	length := len(list)

	// Inverter a ordem das palavras
	for i := 0; i < length/2; i++ {
		list[i], list[length-1-i] = list[length-1-i], list[i]
	}

	// Inverter as letras de cada palavra
	for i := 0; i < length; i++ {
		list[i] = inverterString(list[i])
	}

	fmt.Println(list)
}
