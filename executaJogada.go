package main

import "fmt"

//ExecutaJogada executa a opção escolhida pelo jogador, caso ela esteja disponível
//caso contrário, de forma recursiva, a função é chamada novamente até que a jogada seja válida
func ExecutaJogada(jogador *Jogador, matriz [][]string) bool {
	var jogada string
	fmt.Printf("\nJogador %v - Faça uma jogada - %s", jogador.numero, jogador.caractere)
	fmt.Println("\n\n", RetornaResultadoParcialdaPartida(matriz))
	fmt.Scan(&jogada)
	if VerificaJogadaValida(jogada, matriz) {
		MarcaPosicaoEscolhida(jogada, jogador.caractere, matriz)
	} else {
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		ExecutaJogada(jogador, matriz)
	}
	return VerificaVencedor(matriz) || VerificaVelha(matriz)
}
