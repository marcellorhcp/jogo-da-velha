package main

//IniciaJogo recebe dois ponteiros do tipo Jogador, a matriz que representa o jogo da velha e o map jogadas disponíveis
//e também retorna um valor booleano, um jogador e a matriz final para função InformaResultadoFinalDaPartida
func IniciaJogo(jogador1, jogador2 *Jogador, matriz [][]string, jogadasLivres map[string]int) (bool, int, [][]string) {

	ExecutaJogada(jogador1.player, jogador1.caractere, matriz, jogadasLivres)
	ExecutaJogada(jogador2.player, jogador2.caractere, matriz, jogadasLivres)
	ExecutaJogada(jogador1.player, jogador1.caractere, matriz, jogadasLivres)
	ExecutaJogada(jogador2.player, jogador2.caractere, matriz, jogadasLivres)

	if ExecutaJogada(jogador1.player, jogador1.caractere, matriz, jogadasLivres) {
		return true, jogador1.player, matriz
	}

	if ExecutaJogada(jogador2.player, jogador2.caractere, matriz, jogadasLivres) {
		return true, jogador2.player, matriz
	}

	if ExecutaJogada(jogador1.player, jogador1.caractere, matriz, jogadasLivres) {
		return true, jogador1.player, matriz
	}

	if ExecutaJogada(jogador2.player, jogador2.caractere, matriz, jogadasLivres) {
		return true, jogador2.player, matriz
	}

	if ExecutaJogada(jogador1.player, jogador1.caractere, matriz, jogadasLivres) {
		return true, jogador1.player, matriz
	}
	return false, 0, matriz
}
