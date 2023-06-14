package main

import (
	"fmt"
	"strings"
)

func isPangram(phrase string) bool {
	// Converter a frase para letras minúsculas
	phrase = strings.ToLower(phrase)

	// Criar um mapa para registrar a ocorrência de cada letra
	letterCount := make(map[rune]int)

	// Percorrer cada caractere da frase
	for _, char := range phrase {
		// Verificar se o caractere é uma letra do alfabeto
		if char >= 'a' && char <= 'z' {
			// Incrementar o contador para a letra encontrada
			letterCount[char]++
		}
	}

	// Verificar se todas as letras do alfabeto estão presentes
	return len(letterCount) == 26
}

func main() {
	var phrase string

	// Solicitar ao usuário para inserir uma frase
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&phrase)

	// Verificar se a frase é um pangrama
	if isPangram(phrase) {
		fmt.Println("A frase é um pangrama.")
	} else {
		fmt.Println("A frase não é um pangrama.")
	}
}

