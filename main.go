package main

import "fmt"

type Jogador struct {
	numero    int
	caractere string
}

func main() {
	var jogador1 = novoJogador(1, "")
	var jogador2 = novoJogador(2, "")

	//matriz representa as posições no jogo da velha
	var matriz = [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}

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

	EscolheCaractere(&jogador1, &jogador2)
	InformaResultadoFinalDaPartida(IniciaJogo(&jogador1, &jogador2, matriz))

}
