module localhost.com/bytebank-go

go 1.19

replace contas => ./contas

replace clientes => ./clientes

require (
	clientes v0.0.0-00010101000000-000000000000
	contas v0.0.0-00010101000000-000000000000
)
