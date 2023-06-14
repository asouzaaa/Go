package main

import "fmt"

func main() {
	var inicio, fim int

	fmt.Print("Informe o número inicial do intervalo: ")
	fmt.Scanln(&inicio)

	fmt.Print("Informe o número final do intervalo: ")
	fmt.Scanln(&fim)

	if fim <= inicio {
		fmt.Println("Número final deve ser maior que o número inicial.")
		return
	}

	soma := 0

	for i := inicio; i <= fim; i++ {
		if i%2 == 0 {
			soma += i
		}
	}

	fmt.Println("A soma dos números pares no intervalo é:", soma)
}

