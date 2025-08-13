package conta

import (
	"fmt"
)

type Conta struct {
	Nome  string
	Saldo float64
}

func (c *Conta) Deposit(v float64) {
	if v > 0 {
		c.Saldo += v
	}
}

func (c *Conta) Saque(v float64) {
	if c.Saldo > v {
		c.Saldo -= v
	}
}

func (origem *Conta) Transfer(v float64, destino *Conta) {
	fmt.Println("transferindo")
	if origem.Saldo > v && v > 0 {
		origem.Saque(v)
		destino.Deposit(v)
		return
	}
	fmt.Println("Erro ao transferir")
}
