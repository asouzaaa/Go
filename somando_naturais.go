package main

import "fmt"

func main() {
	inicio := 1
	fim := 100
	soma := 0

	fmt.Printf("Intervalo: %d - %d\n", inicio, fim)

	for i := inicio; i <= fim; i++ {
		soma += i
	}

	fmt.Printf("Soma: %d\n", soma)
}


