package conta_p

import (
	"fmt"
)

type Conta_p struct {
	Nome  string
	Saldo float64
}

func (c *Conta_p) Deposit(v float64) {
	if v > 0 {
		c.Saldo += v
	}
}

func (c *Conta_p) Saque(v float64) {
	if c.Saldo > v {
		c.Saldo -= v
	}
}

func (origem *Conta_p) Transfer(v float64, destino *Conta_p) {
	fmt.Println("transferindo")
	if origem.Saldo > v && v > 0 {
		origem.Saque(v)
		destino.Deposit(v)
		return
	}
	fmt.Println("Erro ao transferir")
}
