package main

import (
	"fmt"
)

func main() {
	var numeros []float64
	var numero float64
	var continuar string

	for {
		fmt.Print("Digite um número: ")
		fmt.Scanln(&numero)

		numeros = append(numeros, numero)

		fmt.Print("Deseja inserir mais números? (S/N): ")
		fmt.Scanln(&continuar)

		if continuar != "S" && continuar != "s" {
			break
		}
	}

	soma := 0.0
	for _, num := range numeros {
		soma += num
	}

	media := soma / float64(len(numeros))

	fmt.Printf("A média dos números informados é: %.2f\n", media)
}

