package main

import "fmt"

func main() {
	var votosA, votosB, votosC, votosNulos, votosEmBranco int

	fmt.Print("Digite a quantidade de votos válidos para o candidato A: ")
	fmt.Scanln(&votosA)

	fmt.Print("Digite a quantidade de votos válidos para o candidato B: ")
	fmt.Scanln(&votosB)

	fmt.Print("Digite a quantidade de votos válidos para o candidato C: ")
	fmt.Scanln(&votosC)

	fmt.Print("Digite a quantidade de votos nulos: ")
	fmt.Scanln(&votosNulos)

	fmt.Print("Digite a quantidade de votos em branco: ")
	fmt.Scanln(&votosEmBranco)

	totalEleitores := votosA + votosB + votosC + votosNulos + votosEmBranco

	percentVotosValidos := float64((votosA+votosB+votosC)*100) / float64(totalEleitores)
	percentVotosA := float64(votosA*100) / float64(totalEleitores)
	percentVotosB := float64(votosB*100) / float64(totalEleitores)
	percentVotosC := float64(votosC*100) / float64(totalEleitores)
	percentVotosNulos := float64(votosNulos*100) / float64(totalEleitores)
	percentVotosEmBranco := float64(votosEmBranco*100) / float64(totalEleitores)

	fmt.Printf("\n--- Resultado da Eleição ---\n")
	fmt.Printf("Número total de eleitores: %d\n", totalEleitores)
	fmt.Printf("Percentual de votos válidos: %.2f%%\n", percentVotosValidos)
	fmt.Printf("Percentual de votos válidos para o candidato A: %.2f%%\n", percentVotosA)
	fmt.Printf("Percentual de votos válidos para o candidato B: %.2f%%\n", percentVotosB)
	fmt.Printf("Percentual de votos válidos para o candidato C: %.2f%%\n", percentVotosC)
	fmt.Printf("Percentual de votos nulos: %.2f%%\n", percentVotosNulos)
	fmt.Printf("Percentual de votos em branco: %.2f%%\n", percentVotosEmBranco)
}

