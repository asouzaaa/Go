package main

import (
	"fmt"
	"sort"
)

func main() {
	// Criar um array de tamanho 12 para armazenar os elementos
	array := make([]int, 12)

	// Ler os elementos do usuário e armazená-los no array
	for i := 0; i < 12; i++ {
		fmt.Printf("Digite o elemento %d: ", i+1)
		fmt.Scan(&array[i])
	}

	// Ordenar o array em ordem decrescente
	sort.Sort(sort.Reverse(sort.IntSlice(array)))

	// Apresentar os elementos ordenados
	fmt.Println("Elementos ordenados em ordem decrescente:")
	for _, elemento := range array {
		fmt.Println(elemento)
	}
}

