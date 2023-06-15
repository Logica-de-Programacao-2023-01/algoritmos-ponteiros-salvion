package main

import "fmt"

//Implemente uma função que receba um ponteiro para uma slice e seu
//tamanho e preencha-o com os n primeiros números primos.

func ehPrimo(i int) bool {
	divisores := 0
	for j := 1; j <= i; j++ {
		if i%j == 0 {
			divisores++
		}
	}
	if divisores == 2 {
		return true
	} else {
		return false
	}
}

func add(ptr *[]int) {
	var n int
	contador := 0
	fmt.Print("Quantos n primos serão adicionados? ")
	fmt.Scan(&n)

	for i := 1; i <= 500; i++ {
		if contador == n {
			break
		}
		if ehPrimo(i) == true {
			*ptr = append(*ptr, i)
			contador++
		}
	}
}

func main() {
	primos := []int{}
	ptr := &primos
	add(ptr)
	fmt.Print(primos)
}
