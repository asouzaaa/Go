package main

import "fmt"

func main() {
	var valorHoraAula float64
	var horasTrabalhadas float64
	var percentualDesconto float64

	fmt.Print("Informe o valor da hora-aula: ")
	fmt.Scanln(&valorHoraAula)

	fmt.Print("Informe o número de horas trabalhadas no mês: ")
	fmt.Scanln(&horasTrabalhadas)

	fmt.Print("Informe o percentual de desconto do INSS: ")
	fmt.Scanln(&percentualDesconto)

	salarioBruto := valorHoraAula * horasTrabalhadas
	desconto := percentualDesconto / 100
	salarioLiquido := salarioBruto * (1 - desconto)

	fmt.Println("Salário bruto: R$", salarioBruto)
	fmt.Println("Salário líquido: R$", salarioLiquido)
}

