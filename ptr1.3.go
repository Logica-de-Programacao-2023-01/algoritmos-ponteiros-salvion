package main

import "fmt"

//Escreva uma função em Go que receba um ponteiro para uma string e
//atualize o valor da string para seu reverso.

func main() {
	s := "Alô?"
	ptr := &s
	fmt.Print(*ptr)
	reverso(ptr)
	fmt.Print(*ptr)
}

func reverso(ptr *string) *string {
	if ptr == nil {
		fmt.Println("Ponteiro nulo.")
	}
	temp := ""
	j := 0
	for i := (len(string(*ptr)) - 1); i >= len(string(*ptr))/2; i-- {
		fmt.Println(len(string(*ptr)))

		temp += string((*ptr)[i])
		j++
	}
	*ptr = temp
	return ptr
}
