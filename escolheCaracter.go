package main

import (
	"fmt"
)

func EscolheCaracter() {
	var escolha string

	fmt.Println("Jogador 1, escolha qual caracter será utilizado")
	fmt.Println("Digite 1 para jogar com X")
	fmt.Println("Digite 2 para jogar com O")
	fmt.Scan(&escolha)

	switch escolha {
	case "1":
		jogador1.caracter = "X"
		jogador2.caracter = "O"
	case "2":
		jogador1.caracter = "O"
		jogador2.caracter = "X"
	default:
		fmt.Println("Parece que você digitou uma opção inválida")
		EscolheCaracter()
	}
}
