package main

import "fmt"

// Crie uma função que receba um ponteiro para uma variável como argumento e modifique o valor da variável dentro da função.
// Certifique-se de que o ponteiro aponte para uma área de memória válida e que a memória seja liberada após o uso.

func modifica(ptr *int) {
	*ptr = 60
}

func main() {
	var num int
	ptr := &num
	modifica(ptr)

	fmt.Print(*ptr)
}
