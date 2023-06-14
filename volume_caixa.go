package main

import "fmt"

func main() {
	var comprimento float64
	var largura float64
	var altura float64

	fmt.Print("Informe o comprimento da caixa: ")
	fmt.Scanln(&comprimento)

	fmt.Print("Informe a largura da caixa: ")
	fmt.Scanln(&largura)

	fmt.Print("Informe a altura da caixa: ")
	fmt.Scanln(&altura)

	volume := comprimento * largura * altura

	fmt.Println("O valor do volume da caixa é:", volume)
}

