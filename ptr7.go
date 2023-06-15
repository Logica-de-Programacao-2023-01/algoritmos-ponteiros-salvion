package main

import "fmt"

//Escreva uma função em Go que receba um ponteiro para um struct Conta com campos saldo e titular,
//e adicione um valor ao saldo da conta.

type Conta struct {
	saldo   float64
	titular string
}

func main() {
	conta1 := Conta{
		saldo:   222.34,
		titular: "Joao",
	}
	ptr := &conta1
	novo_valor := 50.0
	adicionar(ptr, novo_valor)
	fmt.Printf("%.2f", conta1.saldo)
}

func adicionar(c *Conta, valor float64) {
	c.saldo += valor
}
