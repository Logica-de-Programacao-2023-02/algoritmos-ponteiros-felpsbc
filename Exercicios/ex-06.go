package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func atualizarTitulo(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	meuLivro := Livro{
		Titulo: "Jogos Vorazes",
		Autor:  "Anônimo",
	}

	atualizarTitulo(&meuLivro)

	fmt.Printf("Título do livro: %s\n", meuLivro.Titulo)
}
