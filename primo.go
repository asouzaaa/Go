package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	// Verifica se o número é menor que 2, pois números menores que 2 não são primos
	if num < 2 {
		return false
	}

	// Percorre os possíveis divisores do número até a raiz quadrada dele
	// Se encontrar um divisor, o número não é primo
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scanln(&num)

	if isPrime(num) {
		fmt.Printf("%d é um número primo.\n", num)
	} else {
		fmt.Printf("%d não é um número primo.\n", num)
	}
}

