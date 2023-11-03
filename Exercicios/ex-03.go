package main

import "fmt"

func trocar(nome *string) {
	s := []rune(*nome)
	n := len(s)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	*nome = string(s)
}

func main() {
	var nome string
	fmt.Println("Escreva seu nome: ")
	fmt.Scan(&nome)

	trocar(&nome)
	fmt.Println("Nome Invertido: ", nome)

}
