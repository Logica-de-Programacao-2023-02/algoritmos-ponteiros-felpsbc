package main

import "fmt"

type L struct {
	Titulo string
	Autor  string
	Preco  float64
}

func Desconto(livro *L) {
	desconto := 0.10
	livro.Preco = livro.Preco - (livro.Preco * desconto)
}

func main() {
	meuLivro := L{
		Titulo: "Jogos Vorazes",
		Autor:  "Joao Gomes",
		Preco:  380.0,
	}

	Desconto(&meuLivro)

	fmt.Printf("Novo preço do livro: R$%.2f\n", meuLivro.Preco)
}
