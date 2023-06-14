package main

import (
	"fmt"
	"math"
)

func main() {
	var valor1, valor2, valor3 int

	fmt.Print("Informe o primeiro valor inteiro: ")
	fmt.Scanln(&valor1)

	fmt.Print("Informe o segundo valor inteiro: ")
	fmt.Scanln(&valor2)

	fmt.Print("Informe o terceiro valor inteiro: ")
	fmt.Scanln(&valor3)

	somaQuadrados := math.Pow(float64(valor1), 2) + math.Pow(float64(valor2), 2) + math.Pow(float64(valor3), 2)

	fmt.Printf("O valor da soma dos quadrados é: %.2f\n", somaQuadrados)
}

