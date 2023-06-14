package main

import "fmt"

func main() {
	var salario float64
	var reajuste float64

	fmt.Print("Informe o salário mensal: ")
	fmt.Scanln(&salario)

	fmt.Print("Informe o percentual de reajuste: ")
	fmt.Scanln(&reajuste)

	novoSalario := salario + (salario * reajuste / 100)

	fmt.Println("Novo salário:", novoSalario)
}

