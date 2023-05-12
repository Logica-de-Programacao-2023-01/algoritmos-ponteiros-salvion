package main

import "fmt"

//Escreva uma função swap que receba dois ponteiros para int como argumentos e troque
//os valores apontados por eles.

func main() {
	var x int = 12
	var y int = 24
	var ptr1 *int = &x
	var ptr2 *int = &y
	fmt.Println(*ptr1, *ptr2)
	swap(ptr1, ptr2)
	fmt.Println(*ptr1, *ptr2)
}

func swap(ptr1, ptr2 *int) (*int, *int) {
	if ptr1 == nil || ptr2 == nil {
		fmt.Println("Ponteiro é nulo.")
	}[]
	var temp int = *ptr1
	*ptr1 = *ptr2
	*ptr2 = temp
	return ptr1, ptr2
}
