package main

import (
	"fmt"
	"strings"
)

func jokenpo(jogador1, jogador2 string) string {
	switch {
	case jogador1 == jogador2:
		return "Empate"
	case jogador1 == "Tesoura" && (jogador2 == "Papel" || jogador2 == "Lagarto"):
		return "Jogador 1 venceu"
	case jogador1 == "Papel" && (jogador2 == "Pedra" || jogador2 == "Spock"):
		return "Jogador 1 venceu"
	case jogador1 == "Pedra" && (jogador2 == "Lagarto" || jogador2 == "Tesoura"):
		return "Jogador 1 venceu"
	case jogador1 == "Lagarto" && (jogador2 == "Spock" || jogador2 == "Papel"):
		return "Jogador 1 venceu"
	case jogador1 == "Spock" && (jogador2 == "Tesoura" || jogador2 == "Pedra"):
		return "Jogador 1 venceu"
	default:
		return "Jogador 2 venceu"
	}
}

func main() {
	var jogada1, jogada2 string

	fmt.Print("Jogador 1 - Escolha entre Pedra, Papel, Tesoura, Lagarto ou Spock: ")
	fmt.Scanln(&jogada1)

	fmt.Print("Jogador 2 - Escolha entre Pedra, Papel, Tesoura, Lagarto ou Spock: ")
	fmt.Scanln(&jogada2)

	jogada1 = strings.Title(strings.ToLower(jogada1))
	jogada2 = strings.Title(strings.ToLower(jogada2))

	resultado := jokenpo(jogada1, jogada2)
	fmt.Println("Resultado:", resultado)
}

