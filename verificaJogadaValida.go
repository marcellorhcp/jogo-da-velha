package main

//VerificaJogadaValida recebe a jogada escolhida e verifica se a jogada está disponível
func VerificaJogadaValida(jogada string, matriz [][]string) bool {
	a, b := SelecionaPosicao(jogada)
	if a >= 0 && matriz[a][b] == " " {
		return true
	}
	return false
}
