package main

import "fmt"

func main() {
	var cotacaoDolar float64
	var quantidadeReais float64

	fmt.Print("Informe a cotação do dólar: ")
	fmt.Scanln(&cotacaoDolar)

	fmt.Print("Informe a quantidade de reais disponível: ")
	fmt.Scanln(&quantidadeReais)

	valorEmDolar := quantidadeReais / cotacaoDolar

	fmt.Printf("O valor em dólar é: US$ %.2f\n", valorEmDolar)
}

