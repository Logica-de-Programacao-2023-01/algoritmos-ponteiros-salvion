package main

import "fmt"

//Escreva uma função que receba um ponteiro para um inteiro e
//verifique se esse inteiro é par ou ímpar.
//A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.

func main() {
	var x int = 13
	var ptr *int = &x
	fmt.Println("Número:", *ptr)
	im_par(ptr)
	fmt.Println("Atualizado:", *ptr)
}

func im_par(ptr *int) *int {
	if ptr == nil {
		fmt.Println("Ponteiro nulo.")
	}
	if *ptr%2 == 0 {
		*ptr = 0
	} else {
		*ptr = 1
	}
	return ptr
}
