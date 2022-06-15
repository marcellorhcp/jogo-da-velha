package main

import "fmt"

func InformaResultadoFinalDaPartida(resultado bool, jogador int, matriz [][]string) {
	if resultado {
		fmt.Printf("\nParabéns, você venceu Jogador %v!!! :)", jogador)
		fmt.Println("\n\n", RetornaResultadoParcialdaPartida(matriz))
	} else {
		fmt.Println("\nih, deu velha ¯l_( ͡❛ ͜ʖ ͡❛)_/¯")
		fmt.Println("\n\n", RetornaResultadoParcialdaPartida(matriz))
	}
}
