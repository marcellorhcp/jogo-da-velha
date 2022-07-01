package main

import (
	"fmt"

	"github.com/marcellorhcp/jogo-da-velha/model"
	"github.com/marcellorhcp/jogo-da-velha/view"
)

//IniciaJogo recebe dois ponteiros do tipo Jogador e a matriz que representa o jogo da velha
//e também retorna um valor booleano, um jogador e a matriz final para função InformaResultadoFinalDaPartida
func IniciaJogo(jogador1, jogador2 *model.Jogador, matriz [][]string) (bool, int, [][]string) {
	jogadorAtual := jogador1
	proximoJogador := jogador2
	for i := 0; i < 9; i++ {
		if ExecutaJogada(jogadorAtual, matriz) {
			return true, jogadorAtual.Numero, matriz
		}
		jogadorAtual, proximoJogador = proximoJogador, jogadorAtual
	}
	return false, 0, matriz
}

//VerificaJogadaValida recebe a jogada escolhida e verifica se a jogada está disponível
func VerificaJogadaValida(jogada string, matriz [][]string) bool {
	a, b := model.SelecionaPosicao(jogada)
	if a >= 0 && matriz[a][b] == " " {
		return true
	}
	return false
}

//ExecutaJogada executa a opção escolhida pelo jogador caso ela esteja disponível
//caso contrário, de forma recursiva, a função é chamada novamente até que a jogada seja válida
func ExecutaJogada(jogador *model.Jogador, matriz [][]string) bool {
	var jogada string
	fmt.Printf("\nJogador %v - Faça uma jogada - %s", jogador.Numero, jogador.Caractere)
	fmt.Println("\n\n", view.RetornaResultadoParcialdaPartida(matriz))
	fmt.Scan(&jogada)
	view.LimpaTela()
	if VerificaJogadaValida(jogada, matriz) {
		model.MarcaPosicaoEscolhida(jogada, jogador.Caractere, matriz)
	} else {
		view.LimpaTela()
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		ExecutaJogada(jogador, matriz)
	}
	return model.VerificaVencedor(matriz)
}
