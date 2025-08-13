package main

import "fmt"

type Conta struct {
	nome  string
	saldo float64
}

func (c *Conta) deposit(v float64) {
	if v > 0 {
		c.saldo += v
	}
}

func (c *Conta) saque(v float64) {
	if c.saldo > v {
		c.saldo -= v
	}
}

func (origem *Conta) transfer(v float64, destino *Conta) {
	fmt.Println("transferindo")
	if origem.saldo > v && v > 0 {
		origem.saque(v)
		destino.deposit(v)
		return
	}
	fmt.Println("Erro ao transferir")
}

func main() {
	contaMaria := Conta{nome: "maria", saldo: 100}
	contaPedro := Conta{nome: "Pedro", saldo: 100}
	fmt.Println(contaPedro)
	contaPedro.deposit(1000)
	fmt.Println(contaPedro)

	fmt.Println(contaMaria)
	contaMaria.deposit(100)
	fmt.Println(contaMaria)

	contaPedro.transfer(100, &contaMaria)
	fmt.Println(contaPedro)
	fmt.Println(contaMaria)

}
