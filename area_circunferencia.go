package main

import (
	"fmt"
	"math"
)

func main() {
	var raio float64
	pi := 3.14159265

	fmt.Print("Informe o valor do raio: ")
	fmt.Scanln(&raio)

	area := math.Pow(raio, 2) * pi

	fmt.Println("A área da circunferência é:", area)
}

