package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Joao"
	fmt.Println("Hello, ", nome)
	var precoLeite float64 = 2.99
	var precoOvo float64 = 3.99
	var precoPao float64 = 0.99

	fmt.Println("O preço do leite é R$", precoLeite)
	fmt.Println("O preço do ovo é R$", precoOvo)
	fmt.Println("O preço do pão é R$", precoPao)
	fmt.Println("O preço total é R$", precoLeite+precoOvo+precoPao)

	var name = "Maria"
	fmt.Println("O tipo de nome é", reflect.TypeOf(name))
	age := 30
	fmt.Println("O tipo de idade é", reflect.TypeOf(age))
}
