package main

import (
	"fmt"
)

func EscolheCaractere(jogador1, jogador2 *Jogador) (*Jogador, *Jogador) {
	var escolha string

	fmt.Println("Jogador 1, escolha qual caracter será utilizado")
	fmt.Println("Digite 1 para jogar com X")
	fmt.Println("Digite 2 para jogar com O")
	fmt.Scan(&escolha)

	switch escolha {
	case "1":
		jogador1.caractere = "X"
		jogador2.caractere = "O"
	case "2":
		jogador1.caractere = "O"
		jogador2.caractere = "X"
	default:
		fmt.Println("Parece que você digitou uma opção inválida")
		EscolheCaractere(jogador1, jogador2)
	}
	return jogador1, jogador2
}
