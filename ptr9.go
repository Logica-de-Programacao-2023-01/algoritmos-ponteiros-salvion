package main

import "fmt"

//Implemente uma função que receba um ponteiro para uma struct "Livro" com campos "Título",
//"Autor" e "Preço", e que adicione um desconto de 10% no preço do livro.

type Livro struct {
	titulo string
	autor  string
	preco  float64
}

func desconto(li *Livro) {
	li.preco -= 0.1 * li.preco
}

func main() {
	livro := Livro{
		titulo: "Perspicaz",
		autor:  "Maria Dona",
		preco:  100.00,
	}
	ptr := &livro
	desconto(ptr)
	fmt.Printf("%.2f", livro.preco)
}
