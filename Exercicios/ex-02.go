package main

import "fmt"

func decide(num *int, n int) {
	if n < 0 {
		fmt.Println("O número deve ser maior que 0")
		return
	}
	if n%2 == 1 {
		fmt.Println("O número escolhido é Impar")
		*num = 1
	} else {
		fmt.Println("O número escolhido é Par")
		*num = 0
	}
}

func main() {
	var resultado int
	var n int
	fmt.Println("Escreva um número para descobrir se ele é (1)impar ou (0)par: ")
	fmt.Scan(&n)

	decide(&resultado, n)

	fmt.Printf("O valor atualizado do número é: %d\n", resultado)

}
