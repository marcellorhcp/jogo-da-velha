package main

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

	InstrucoesDoJogo()

	EscolheCaractere(&jogador1, &jogador2)
	InformaResultadoFinalDaPartida(IniciaJogo(&jogador1, &jogador2, matriz))

}
