package view

import "fmt"

func InstrucoesDoJogo() {
	fmt.Printf(`
	###############################################################
	#                   Bem-vindo ao Jogo da Velha!               #         
	# O jogo funciona da seguinte forma:                          #
	# Cada número escolhido representa uma casa no jogo da velha  #
	# A ordem das casas está de acordo com o teclado numérico     #
	#                          Boa sorte!                         #
	###############################################################

	                        +---+---+---+
	                        | 7 | 8 | 9 |
	                        +---+---+---+
	                        | 4 | 5 | 6 |
	                        +---+---+---+
	                        | 1 | 2 | 3 |
	                        +---+---+---+

	`)
	EnterParaContinuar()
}
