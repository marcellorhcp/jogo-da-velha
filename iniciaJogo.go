package main

//IniciaJogo recebe dois ponteiros do tipo Jogador, a matriz que representa o jogo da velha e o map jogadas disponíveis
//e também retorna um valor booleano, um jogador e a matriz final para função InformaResultadoFinalDaPartida
func IniciaJogo(jogador1, jogador2 *Jogador, matriz [][]string) (bool, int, [][]string) {

	ExecutaJogada(jogador1, matriz)
	ExecutaJogada(jogador2, matriz)
	ExecutaJogada(jogador1, matriz)
	ExecutaJogada(jogador2, matriz)

	if ExecutaJogada(jogador1, matriz) {
		return true, jogador1.numero, matriz
	}

	if ExecutaJogada(jogador2, matriz) {
		return true, jogador2.numero, matriz
	}

	if ExecutaJogada(jogador1, matriz) {
		return true, jogador1.numero, matriz
	}
	if VerificaVelha(matriz) {
		return false, 0, matriz
	}
	if ExecutaJogada(jogador2, matriz) {
		return true, jogador2.numero, matriz
	}
	if VerificaVelha(matriz) {
		return false, 0, matriz
	}
	if ExecutaJogada(jogador1, matriz) {
		return true, jogador1.numero, matriz
	}
	return false, 0, matriz
}
