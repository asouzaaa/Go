package main

import (
	"fmt"
)

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

	fizzBuzz(inicio, fim)
}

func fizzBuzz(inicio, fim int) {
	for i := inicio; i <= fim; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

