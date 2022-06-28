package main

//IniciaJogo recebe dois ponteiros do tipo Jogador e a matriz que representa o jogo da velha
//e também retorna um valor booleano, um jogador e a matriz final para função InformaResultadoFinalDaPartida
func IniciaJogo(jogador1, jogador2 *Jogador, matriz [][]string) (bool, int, [][]string) {
	jogadorAtual := jogador1
	proximoJogador := jogador2
	for i := 0; i < 9; i++ {
		if ExecutaJogada(jogadorAtual, matriz) {
			return true, jogadorAtual.numero, matriz
		}
		jogadorAtual, proximoJogador = proximoJogador, jogadorAtual
	}
	return false, 0, matriz
}
