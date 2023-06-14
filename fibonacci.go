package main

import "fmt"

func fibonacci(n int) []int {
	// Inicializar a sequência de Fibonacci com os primeiros dois números: 0 e 1
	sequence := []int{0, 1}

	// Gerar os próximos números da sequência até alcançar ou ultrapassar o número informado
	for i := 2; sequence[i-1] <= n; i++ {
		next := sequence[i-1] + sequence[i-2]
		sequence = append(sequence, next)
	}

	return sequence
}

func main() {
	var num int

	// Solicitar ao usuário para inserir o número
	fmt.Print("Digite um número: ")
	fmt.Scanln(&num)

	// Obter a sequência de Fibonacci até o número informado
	result := fibonacci(num)

	// Imprimir a sequência de Fibonacci
	fmt.Println("Sequência de Fibonacci:")
	for _, n := range result {
		fmt.Print(n, " ")
	}
	fmt.Println()
}

