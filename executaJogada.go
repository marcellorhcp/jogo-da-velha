package main

import (
	"fmt"

	"github.com/marcellorhcp/jogo-da-velha/model"
	"github.com/marcellorhcp/jogo-da-velha/view"
)

//ExecutaJogada executa a opção escolhida pelo jogador caso ela esteja disponível
//caso contrário, de forma recursiva, a função é chamada novamente até que a jogada seja válida
func ExecutaJogada(jogador *model.Jogador, matriz [][]string) bool {
	var jogada string
	fmt.Printf("\nJogador %v - Faça uma jogada - %s", jogador.Numero, jogador.Caractere)
	fmt.Println("\n\n", view.RetornaResultadoParcialdaPartida(matriz))
	fmt.Scan(&jogada)
	view.LimpaTela()
	if VerificaJogadaValida(jogada, matriz) {
		MarcaPosicaoEscolhida(jogada, jogador.Caractere, matriz)
	} else {
		view.LimpaTela()
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		ExecutaJogada(jogador, matriz)
	}
	return VerificaVencedor(matriz)
}
