package main

import "fmt"

func ExecutaJogada(numeroDoJogador int, opcaoDeJogada string, matriz [][]string) bool {
	var jogada string

	fmt.Printf("\nJogador %v - Faça uma jogada - %s", numeroDoJogador, opcaoDeJogada)
	fmt.Println("\n\n", RetornaResultadoParcialdaPartida(matriz))
	fmt.Scan(&jogada)
	if VerificaJogadaValida(jogada) {
		MarcaPosicaoEscolhida(jogada, opcaoDeJogada, matriz)
	} else {
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		ExecutaJogada(numeroDoJogador, opcaoDeJogada, matriz)
	}
	return VerificaVencedor(matriz)
}
