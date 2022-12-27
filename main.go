package main

import (
	"clientes"
	"contas"
	"fmt"
)

type Conta interface {
	Sacar(valor float64) (string, float64)
}

func pagarBoleto(conta Conta, valor float64) {
	conta.Sacar(valor)
}

func main() {
	titular1 := clientes.Titular{Nome: "Elton", CPF: "123.456.789-10", Profissao: "Desenvolvedor GO"}
	cc := contas.ContaCorrente{Titular: titular1, NumeroAgencia: 123, NumeroConta: 654321}

	titular2 := clientes.Titular{Nome: "Matheus Guimaraes", CPF: "321.654.987-10", Profissao: "Desenvolvedor JAVA"}
	cp := contas.ContaPoupanca{Titular: titular2, NumeroAgencia: 321, NumeroConta: 123456}

	cc.Depositar(100)
	pagarBoleto(&cc, 60)

	cp.Depositar(200)
	pagarBoleto(&cp, 120)

	fmt.Println(cc)
	fmt.Println(cp)
}
