package main

import "fmt"

func TwoFer(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

func main() {
	var name string
	fmt.Print("Informe um nome: ")
	fmt.Scanln(&name)
	result := TwoFer(name)
	fmt.Println(result)
}

