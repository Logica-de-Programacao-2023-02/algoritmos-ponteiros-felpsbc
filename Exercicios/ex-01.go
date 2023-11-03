package main

import "fmt"

func numero(num *int, n int) {
	if n < 0 {
		fmt.Println("O número deve ser positivo")
		return
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*num = sum
}

func main() {
	var resultado int
	n := 5

	numero(&resultado, n)

	fmt.Printf("A soma dos primeiros %d números naturais é: %d\n", n, resultado)
}
