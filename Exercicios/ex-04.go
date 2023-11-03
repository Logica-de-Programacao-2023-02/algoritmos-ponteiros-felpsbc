package main

import "fmt"

func soma(num *int) {
	numero := *num
	if numero < 0 {
		numero = -numero
	}
	ultimo := numero % 10
	numero = numero / 10
	penultimo := numero % 10

	soma := ultimo + penultimo

	*num = soma
}

func main() {
	var valor int
	fmt.Println("Escreva um valor")
	fmt.Scan(&valor)

	soma(&valor)

	fmt.Printf("Valor da soma: %d\n", valor)
}
