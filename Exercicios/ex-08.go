package main

import "fmt"

func modificar(valor *int) {
	novoValor := 120
	*valor = novoValor
}

func main() {
	num := 50

	fmt.Printf("Antes da função: %d\n", num)
	modificar(&num)
	fmt.Printf("Depois da função: %d\n", num)
}
