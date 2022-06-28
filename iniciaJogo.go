package main

//IniciaJogo recebe dois ponteiros do tipo Jogador e a matriz que representa o jogo da velha
//e também retorna um valor booleano, um jogador e a matriz final para função InformaResultadoFinalDaPartida
func IniciaJogo(jogador1, jogador2 *Jogador, matriz [][]string) (bool, int, [][]string) {
	for jogador1.quantidadeJogadas < 5 {
		if ExecutaJogada(jogador1, matriz) {
			return true, jogador1.numero, matriz
		}
		// if VerificaVelha(matriz) {
		// 	return false, 0, matriz
		// }
		if jogador1.quantidadeJogadas < 4 {
			if ExecutaJogada(jogador2, matriz) {
				return true, jogador2.numero, matriz
			}
			// if VerificaVelha(matriz) {
			// 	return false, 0, matriz
			// }
		}
		jogador1.quantidadeJogadas++
		jogador2.quantidadeJogadas++
	}
	return false, 0, matriz
}
