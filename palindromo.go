package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// Remover espaços em branco e converter para minúsculas
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	// Comparar a sequência original com a sequência invertida
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false // Não é um palíndromo
		}
	}

	return true // É um palíndromo
}

func main() {
	var entrada string
	fmt.Print("Digite uma palavra, frase ou número: ")
	fmt.Scanln(&entrada)

	fmt.Printf("Entrada: %s\n", entrada)

	if isPalindrome(entrada) {
		fmt.Println("Saída: true")
		fmt.Println("Explicação: É um palíndromo")
	} else {
		fmt.Println("Saída: false")
		fmt.Println("Explicação: Não é um palíndromo")
	}
}



