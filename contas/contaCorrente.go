package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (cc *ContaCorrente) Tranferir(valor float64, conta *ContaCorrente) string {
	if valor <= 0 {
		return "Valor de tranferência inválido"
	}

	temSaldo := cc.Saldo >= valor

	if temSaldo {
		conta.depositar(valor)
		cc.Sacar(valor)
		return "Tranferência com sucesso"
	} else {
		return "Conta sem Saldo"
	}
}

func (cc *ContaCorrente) depositar(valor float64) {
	if valor > 0 {
		cc.Saldo += valor
	}
}

func (cc *ContaCorrente) Sacar(valor float64) (string, float64) {
	var message string

	if valor <= 0 {
		message = "Valor de saque inválido"
	}

	podeSacar := valor <= cc.Saldo

	if podeSacar {
		cc.Saldo -= valor
		message = "Saque realizado com sucesso"

	} else {
		message = "Saque insuficiente"
	}

	return message, cc.Saldo
}
