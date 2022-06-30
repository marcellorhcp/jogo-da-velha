package main

import (
	"fmt"

	"github.com/marcellorhcp/jogo-da-velha/view"
)

//InformaResultadoFinalDaPartida recebe o retorno da função IniciaJogo e mostra na tela a mensagem final do resultado da partida
func InformaResultadoFinalDaPartida(resultado bool, jogador int, matriz [][]string) {
	if resultado {
		fmt.Printf("\nParabéns, você venceu Jogador %v!!! :)", jogador)
		fmt.Println("\n\n", view.RetornaResultadoParcialdaPartida(matriz))
	} else {
		fmt.Println("\nih, deu velha ¯l_( ͡❛ ͜ʖ ͡❛)_/¯")
		fmt.Println("\n\n", view.RetornaResultadoParcialdaPartida(matriz))
	}
}
