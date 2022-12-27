package main

import (
	"fmt"

	"./contas"
)

func main() {
	cc := contas.ContaCorrente{Titular: "André", NumeroAgencia: 123, NumeroConta: 654321, Saldo: 500.0}
	cc1 := contas.ContaCorrente{Titular: "Elton", NumeroAgencia: 321, NumeroConta: 123456, Saldo: 0.0}

	fmt.Println("cc1: ", cc1)

	cc.Tranferir(100, &cc)

	fmt.Println("SALDO: ", cc.Saldo)
	fmt.Println("cc1: ", cc1)
}
