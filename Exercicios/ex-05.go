package main

import (
	"fmt"
	"math"
)

func atualizarComMediaPi(valor *float64) {
	valorAtual := *valor

	media := (valorAtual + math.Pi) / 2.0

	*valor = media
}

func main() {
	meuNumero := 5.70

	atualizarComMediaPi(&meuNumero)

	fmt.Printf("Novo valor: %.2f\n", meuNumero)
}
