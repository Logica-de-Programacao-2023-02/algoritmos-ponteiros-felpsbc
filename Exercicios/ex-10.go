package main

import "fmt"

func Primo(numero int) bool {
	if numero <= 1 {
		return false
	}
	if numero <= 3 {
		return true
	}
	if numero%2 == 0 || numero%3 == 0 {
		return false
	}
	i := 5
	for i*i <= numero {
		if numero%i == 0 || numero%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func ColocarPrimos(slice *[]int, tamanho int) {
	primos := make([]int, 0, tamanho)
	numero := 2

	for len(primos) < tamanho {
		if Primo(numero) {
			primos = append(primos, numero)
		}
		numero++
	}

	*slice = primos
}

func main() {
	tamanho := 10
	numerosPrimos := make([]int, 0, tamanho)

	ColocarPrimos(&numerosPrimos, tamanho)

	fmt.Printf("Números primos: %v\n", numerosPrimos)
}
