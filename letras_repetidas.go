package main

import (
	"fmt"
	"strings"
)

func main() {
	var texto string

	// Solicitar ao usuário para inserir o texto
	fmt.Print("Digite um texto: ")
	fmt.Scanln(&texto)

	// Remover espaços em branco do texto e converter para letras minúsculas
	texto = strings.ToLower(strings.ReplaceAll(texto, " ", ""))

	// Criar um mapa para contar a ocorrência de cada letra
	ocorrencias := make(map[rune]int)

	// Contar as ocorrências de cada letra no texto
	for _, char := range texto {
		ocorrencias[char]++
	}

	// Encontrar a primeira letra não repetida
	for _, char := range texto {
		if ocorrencias[char] == 1 {
			fmt.Printf("A primeira letra não repetida é: %c\n", char)
			return
		}
	}

	fmt.Println("Não existem letras que não se repetem no texto informado.")
}

