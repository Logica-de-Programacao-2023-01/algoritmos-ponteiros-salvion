package main

import "fmt"

//Escreva uma função que receba um ponteiro para uma struct que contém informações
//de um produto (nome, preço e quantidade em estoque). A função deve atualizar o preço
//desse produto para um novo valor fornecido como argumento.

func main() {
	s := Produto{
		nome:  "Sorvete Kibon",
		tipo:  "Congelados",
		preco: 24.99,
	}
	fmt.Println("O antigo preço era:", s.preco)
	novoPreco(&s, 27.99)
	fmt.Println("O novo preço é:", s.preco)
}

func novoPreco(s *Produto, novo_preco float64) float64 {
	s.preco = novo_preco
	return s.preco
}

type Produto struct {
	nome  string
	tipo  string
	preco float64
}
