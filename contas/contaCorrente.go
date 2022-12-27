package contas

import (
	"clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (cc *ContaCorrente) Tranferir(valor float64, conta *ContaCorrente) string {
	if valor <= 0 {
		return "Valor de tranferência inválido"
	}

	temsaldo := cc.saldo >= valor

	if temsaldo {
		conta.Depositar(valor)
		cc.Sacar(valor)
		return "Tranferência com sucesso"
	} else {
		return "Conta sem saldo"
	}
}

func (cc *ContaCorrente) Depositar(valor float64) {
	if valor > 0 {
		cc.saldo += valor
	}
}

func (cc *ContaCorrente) Sacar(valor float64) (string, float64) {
	var message string

	if valor <= 0 {
		message = "Valor de saque inválido"
	}

	podeSacar := valor <= cc.saldo

	if podeSacar {
		cc.saldo -= valor
		message = "Saque realizado com sucesso"

	} else {
		message = "Saque insuficiente"
	}

	return message, cc.saldo
}

func (cc *ContaCorrente) ObterSaldo() float64 {
	return cc.saldo
}
