package main

import (
	"github.com/marcellorhcp/jogo-da-velha/model"
	"github.com/marcellorhcp/jogo-da-velha/view"
)

func main() {

	var jogador1 = model.NovoJogador(1, "")
	var jogador2 = model.NovoJogador(2, "")

	//matriz representa as posições no jogo da velha
	var matriz = [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}

	view.InstrucoesDoJogo()

	view.EscolheCaractere(&jogador1, &jogador2)
	view.InformaResultadoFinalDaPartida(IniciaJogo(&jogador1, &jogador2, matriz))

}
