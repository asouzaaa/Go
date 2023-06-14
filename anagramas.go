package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str1, str2 string

	// Solicitar ao usuário para inserir as strings
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	// Remover espaços em branco das strings e converter para letras minúsculas
	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	// Verificar se as strings têm o mesmo tamanho
	if len(str1) != len(str2) {
		fmt.Println("As strings não são anagramas.")
		return
	}

	// Ordenar as letras das strings
	str1Sorted := sortString(str1)
	str2Sorted := sortString(str2)

	// Verificar se as strings ordenadas são iguais
	if str1Sorted == str2Sorted {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

// Função para ordenar os caracteres de uma string
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

