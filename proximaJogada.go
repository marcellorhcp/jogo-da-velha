package main

import "fmt"

func ProximaJogada(num int) {
	jogador = num
	var opcao string = "X"
	if jogador == 2 {
		opcao = "0"
	}
	fmt.Printf("\nJogador %v - Faça uma jogada - %s", jogador, opcao)
	fmt.Println("\n", RetornaResultadoParcialdaPartida())
	fmt.Scan(&jogada)
	if VerificaJogadaValida(jogada) {
		MarcaPosicaoEscolhida(jogada, opcao)
	} else {
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		ProximaJogada(jogador)
	}
}