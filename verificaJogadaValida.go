package main

//VerificaJogadaValida recebe a jogada escolhida e através de um map, verifica se a jogada está disponível
func VerificaJogadaValida(jogada string, matriz [][]string) bool {
	a, b := SelecionaPosicao(jogada)
	if matriz[a][b] == " " {
		return true
	}
	return false
}
