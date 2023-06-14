package main

import (
	"fmt"
	"strings"
)

func breakLines(sentence string, columns int) string {
	words := strings.Split(sentence, " ")
	lines := make([]string, 0)
	currentLine := ""

	for _, word := range words {
		// Verificar se a palavra cabe na linha atual
		if len(currentLine)+len(word)+1 <= columns {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		} else {
			// Adicionar a linha atual às linhas resultantes
			lines = append(lines, currentLine)

			// Iniciar uma nova linha com a palavra atual
			currentLine = word
		}
	}

	// Adicionar a última linha à lista de linhas resultantes
	lines = append(lines, currentLine)

	// Juntar as linhas em uma única string com quebras de linha
	result := strings.Join(lines, "\n")

	return result
}

func main() {
	var sentence string
	var columns int

	// Solicitar ao usuário a frase
	fmt.Print("Digite a frase: ")
	fmt.Scanln(&sentence)

	// Solicitar ao usuário a quantidade de colunas
	fmt.Print("Digite a quantidade de colunas: ")
	fmt.Scanln(&columns)

	// Quebrar as linhas da frase
	result := breakLines(sentence, columns)

	// Imprimir o resultado
	fmt.Println(result)
}

