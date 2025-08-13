package main

import (
	"banco_pkg/calculadora"
	"banco_pkg/conta"
	"banco_pkg/conta_p"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Saque(valorDoBoleto)
}

type verificarConta interface {
	Saque(valor float64)
}

func main() {
	contaMaria := conta.Conta{Nome: "maria", Saldo: 100}
	contaPedro := conta.Conta{Nome: "Pedro", Saldo: 100}
	fmt.Println(contaPedro)
	contaPedro.Deposit(1000)
	fmt.Println(contaPedro)

	fmt.Println(contaMaria)
	contaMaria.Deposit(100)
	fmt.Println(contaMaria)

	contaPedro.Transfer(100, &contaMaria)
	fmt.Println(contaPedro)
	fmt.Println(contaMaria)

	fmt.Println("saldo do banco é:", calculadora.Sum(contaMaria.Saldo, contaPedro.Saldo))

	contaSergio := conta_p.Conta_p{Nome: "Sergio", Saldo: 100.0}

	PagarBoleto(&contaPedro, 50)
	PagarBoleto(&contaSergio, 50)

	fmt.Println(contaPedro)
	fmt.Println(contaSergio)

}
