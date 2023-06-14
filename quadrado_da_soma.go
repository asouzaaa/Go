package main

import "fmt"

func main() {
	var valor1, valor2, valor3 int

	fmt.Print("Informe o primeiro valor inteiro: ")
	fmt.Scanln(&valor1)

	fmt.Print("Informe o segundo valor inteiro: ")
	fmt.Scanln(&valor2)

	fmt.Print("Informe o terceiro valor inteiro: ")
	fmt.Scanln(&valor3)

	soma := valor1 + valor2 + valor3
	quadradoDaSoma := soma * soma

	fmt.Printf("O valor do quadrado da soma é: %d\n", quadradoDaSoma)
}

