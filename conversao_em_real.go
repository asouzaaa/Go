package main

import "fmt"

func main() {
	var cotacaoDolar float64
	var quantidadeDolar float64

	fmt.Print("Informe a cotação do dólar: ")
	fmt.Scanln(&cotacaoDolar)

	fmt.Print("Informe a quantidade de dólares disponível: ")
	fmt.Scanln(&quantidadeDolar)

	valorEmReal := cotacaoDolar * quantidadeDolar

	fmt.Printf("O valor em reais é: R$ %.2f\n", valorEmReal)
}

