package main

import "fmt"

type jogador struct {
	caractere string
	player    int
	//nome string
}

var jogador1 = jogador{
	player:    1,
	caractere: "",
}
var jogador2 = jogador{
	player:    2,
	caractere: "",
}

//var jogador int
var andamentoDoJogo string
var jogada string
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

func main() {
	// var nome string
	// fmt.Println("Digite o nome do Jogador 1")
	// fmt.Scan(&nome)

	fmt.Printf(`
	# Bem-vindo ao Jogo da Velha
	# O jogo funciona da seguinte forma:
	# Cada número escolhido representa uma casa no jogo da velha
	# A ordem das casas está de acordo com o teclado numérico
	# Boa sorte
	
		| 7 | 8 | 9 |
		------------
		| 4 | 5 | 6 |
		------------
		| 1 | 2 | 3 |

`)
	EscolheCaractere()
	InformaResultadoFinalDaPartida(IniciaJogo())
}
