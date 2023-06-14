package main

import "fmt"

func romanToInt(s string) int {
	// Mapeamento dos símbolos romanos e seus respectivos valores
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0

	// Percorre a string de trás para frente
	for i := len(s) - 1; i >= 0; i-- {
		// Obtém o valor do símbolo atual
		value := romanMap[s[i]]

		// Se o valor atual for menor que o valor anterior,
		// subtrai esse valor do resultado
		if value < prevValue {
			result -= value
		} else {
			// Caso contrário, adiciona o valor ao resultado
			result += value
			prevValue = value
		}
	}

	return result
}

func main() {
	var s string
	fmt.Print("Digite um numeral romano: ")
	fmt.Scanln(&s)

	result := romanToInt(s)
	fmt.Printf("O numeral romano %s é equivalente a %d\n", s, result)
}

