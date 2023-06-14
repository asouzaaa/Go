package main

import "fmt"

func main() {
	var A, B int

	fmt.Print("Informe o valor de A: ")
	fmt.Scanln(&A)

	fmt.Print("Informe o valor de B: ")
	fmt.Scanln(&B)

	// Efetua a troca dos valores
	A, B = B, A

	fmt.Println("Valores após a troca:")
	fmt.Println("A:", A)
	fmt.Println("B:", B)
}

