package main

//Escreva uma função que receba um ponteiro para um booleano e altere o valor desse
//booleano para o seu inverso.

import "fmt"

func main() {
	var um bool = true
	var ptr1 *bool = &um
	fmt.Println(*ptr1)
	inverso(ptr1)
	fmt.Println(*ptr1)
}

func inverso(ptr1 *bool) *bool {
	if ptr1 == nil {
		fmt.Println("Ponteiro nulo.")
	}
	if *ptr1 == true {
		*ptr1 = false
	} else {
		*ptr1 = true
	}
	return ptr1
}
