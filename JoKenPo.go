package main

import (
	"fmt"
	"strings"
)

func jokenpo(jogador1, jogador2 string) string {
	if jogador1 == jogador2 {
		return "Empate"
	}

	if jogador1 == "Pedra" && jogador2 == "Tesoura" ||
		jogador1 == "Tesoura" && jogador2 == "Papel" ||
		jogador1 == "Papel" && jogador2 == "Pedra" {
		return "Jogador 1 venceu"
	}

	return "Jogador 2 venceu"
}

func main() {
	var jogada1, jogada2 string

	fmt.Print("Jogador 1 - Escolha entre Pedra, Papel ou Tesoura: ")
	fmt.Scanln(&jogada1)

	fmt.Print("Jogador 2 - Escolha entre Pedra, Papel ou Tesoura: ")
	fmt.Scanln(&jogada2)

	jogada1 = strings.Title(strings.ToLower(jogada1))
	jogada2 = strings.Title(strings.ToLower(jogada2))

	resultado := jokenpo(jogada1, jogada2)
	fmt.Println("Resultado:", resultado)
}

