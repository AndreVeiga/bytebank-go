package contas

import (
	"clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (cc *ContaPoupanca) Depositar(valor float64) {
	if valor > 0 {
		cc.saldo += valor
	}
}

func (cp *ContaPoupanca) Sacar(valor float64) (string, float64) {
	var message string

	if valor <= 0 {
		message = "Valor de saque inválido"
	}

	podeSacar := valor <= cp.saldo

	if podeSacar {
		cp.saldo -= valor
		message = "Saque realizado com sucesso"

	} else {
		message = "Saque insuficiente"
	}

	return message, cp.saldo
}
