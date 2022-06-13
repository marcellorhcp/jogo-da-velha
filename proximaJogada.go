package main

import "fmt"

func ProximaJogada(numeroDoJogador int, opcaoDeJogada string) {
	fmt.Printf("\nJogador %v - Faça uma jogada - %s", numeroDoJogador, opcaoDeJogada)
	fmt.Println("\n\n", RetornaResultadoParcialdaPartida())
	fmt.Scan(&jogada)
	if VerificaJogadaValida(jogada) {
		MarcaPosicaoEscolhida(jogada, opcaoDeJogada)

	} else {
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		ProximaJogada(numeroDoJogador, opcaoDeJogada)
	}
}
