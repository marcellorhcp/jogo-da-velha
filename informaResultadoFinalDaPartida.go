package main

import "fmt"

func InformaResultadoFinalDaPartida() {
	if vencedor {
		fmt.Printf("\nParabéns, você venceu Jogador %v!!! :)", jogador)
		fmt.Println("\n\n", RetornaResultadoParcialdaPartida())
	} else {
		fmt.Println("\nih, deu velha ¯l_( ͡❛ ͜ʖ ͡❛)_/¯")
		fmt.Println("\n\n", RetornaResultadoParcialdaPartida())
	}
}
