package main

import (
	"fmt"
)

func calculateChange(total, paid float64) map[string]int {
	// Definir os valores das cédulas e moedas
	notes := []float64{100, 50, 10, 5, 1}
	coins := []float64{0.50, 0.10, 0.05, 0.01}

	// Calcular o troco a ser fornecido
	change := paid - total

	// Inicializar o mapa para armazenar a quantidade de cédulas e moedas
	changeCount := make(map[string]int)

	// Calcular o número de cédulas de troco
	for _, note := range notes {
		count := int(change / note)
		changeCount[fmt.Sprintf("R$%.2f", note)] = count
		change -= float64(count) * note
	}

	// Calcular o número de moedas de troco
	for _, coin := range coins {
		count := int(change / coin)
		changeCount[fmt.Sprintf("R$%.2f", coin)] = count
		change -= float64(count) * coin
	}

	return changeCount
}

func main() {
	var total, paid float64

	// Solicitar ao usuário o valor total a ser pago
	fmt.Print("Valor total a ser pago: ")
	fmt.Scanln(&total)

	// Solicitar ao usuário o valor efetivamente pago
	fmt.Print("Valor efetivamente pago: ")
	fmt.Scanln(&paid)

	// Calcular o troco
	change := calculateChange(total, paid)

	// Imprimir o número de cédulas e moedas de troco
	fmt.Println("Troco:")
	for currency, count := range change {
		fmt.Printf("%s: %d\n", currency, count)
	}
}


