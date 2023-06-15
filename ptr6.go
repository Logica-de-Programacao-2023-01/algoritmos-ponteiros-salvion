package main

import "fmt"

//Escreva uma função em Go que receba um ponteiro para um struct Livro com campos título e autor,
//e altere o título do livro para "Desconhecido" se o autor for "Anônimo".

type Livro struct {
	titulo string
	autor  string
}

func main() {
	livro := Livro{
		titulo: "livro x",
		autor:  "Anônimo",
	}
	check(&livro)
	fmt.Print("Mudança:", livro)
}

func check(l *Livro) {
	if l.autor == "Anônimo" {
		l.titulo = "Desconhecido"
	}
}
