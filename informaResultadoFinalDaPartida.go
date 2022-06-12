package main

import "fmt"

func InformaResultadoFinalDaPartida(resultado bool) {
	if resultado {
		fmt.Printf("\nParabéns, você venceu Jogador %v!!! :)", jogador)
		fmt.Println("\n\n", RetornaResultadoParcialdaPartida())
	} else {
		fmt.Println("\nih, deu velha ¯l_( ͡❛ ͜ʖ ͡❛)_/¯")
		fmt.Println("\n\n", RetornaResultadoParcialdaPartida())
	}
}
