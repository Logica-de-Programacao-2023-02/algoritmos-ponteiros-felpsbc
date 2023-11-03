package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarSaldo(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	minhaConta := Conta{
		Saldo:   1890.50,
		Titular: "Felipe",
	}

	adicionarSaldo(&minhaConta, 1000.0)

	fmt.Printf("Novo saldo: %.2f\n", minhaConta.Saldo)
}
