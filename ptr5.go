package main

import "fmt"

// Escreva uma função em Go que receba um ponteiro para uma variável float64 e
// atualize o valor da variável para a média aritmética entre o seu valor atual e o valor da constante Pi.
func main() {
	num := 6.9
	var ptr *float64 = &num
	*ptr = (num + 3.1415) / 2
	fmt.Print(*ptr)
}
