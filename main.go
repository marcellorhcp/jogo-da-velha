package main

import "fmt"

var jogador int
var andamentoDoJogo string
var jogada string
var vencedor bool = false
var matriz = [][]string{
	{" ", " ", " "},
	{" ", " ", " "},
	{" ", " ", " "},
}
var jogadasLivres = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func jogadaJogador1() {
	jogador = 1
	fmt.Println("\nJogador 1 - Faça uma jogada - X")
	fmt.Println("\n", RetornaResultadoParcialdaPartida())
	fmt.Scan(&jogada)
	if VerificaJogadaValida(jogada) {
		MarcaPosicaoEscolhida(jogada, jogador)
	} else {
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		jogadaJogador1()
	}
}
func jogadaJogador2() {
	jogador = 2
	fmt.Println("\nJogador 2 - Faça uma jogada - O")
	fmt.Println("\n", RetornaResultadoParcialdaPartida())
	fmt.Scan(&jogada)
	if VerificaJogadaValida(jogada) {
		MarcaPosicaoEscolhida(jogada, jogador)
	} else {
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		jogadaJogador2()
	}
}

func main() {
	// var nome string
	// fmt.Println("Digite o nome do Jogador 1")
	// fmt.Scan(&nome)

	fmt.Printf(`
	# Bem-vindo ao Jogo da Velha
	# O jogo funciona da seguinte forma:
	# Cada número escolhido representa uma casa no jogo da velha
	# A ordem das casas está de acordo com o teclado númerico
	# Boa sorte
	
		| 7 | 8 | 9 |
		------------
		| 4 | 5 | 6 |
		------------
		| 1 | 2 | 3 |

`)
	IniciaJogo()
	InformaResultadoFinalDaPartida()
}
