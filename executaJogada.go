package main

import "fmt"

//ExecutaJogada executa a opção escolhida pelo jogador, caso ela esteja disponível
//caso contrário, de forma recursiva, a função é chamada novamente até que a jogada seja válida
func ExecutaJogada(numeroDoJogador int, opcaoDeJogada string, matriz [][]string, jogadasLivres map[string]int) bool {
	var jogada string

	fmt.Printf("\nJogador %v - Faça uma jogada - %s", numeroDoJogador, opcaoDeJogada)
	fmt.Println("\n\n", RetornaResultadoParcialdaPartida(matriz))
	fmt.Scan(&jogada)
	if VerificaJogadaValida(jogada, jogadasLivres) {
		MarcaPosicaoEscolhida(jogada, opcaoDeJogada, matriz)
	} else {
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		ExecutaJogada(numeroDoJogador, opcaoDeJogada, matriz, jogadasLivres)
	}
	return VerificaVencedor(matriz)
}
