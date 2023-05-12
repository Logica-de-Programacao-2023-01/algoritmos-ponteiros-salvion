package main

//Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n.
//A função deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.

import "fmt"

func main() {
	var x int = 12
	var ptr *int = &x
	var n int = 10
	fmt.Println("Valor anterior:", *ptr)
	SomaN(ptr, n)
	fmt.Println("Valor atual:", *ptr)
}

func SomaN(ptr *int, n int) *int {
	if ptr == nil {
		fmt.Println("Ponteiro nulo.")
	}
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	*ptr = soma
	return ptr
}
