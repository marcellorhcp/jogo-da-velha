package main

//VerificaJogadaValida recebe a jogada escolhida e através de um map, verifica se a jogada está disponível
func VerificaJogadaValida(jogadaEscolhida string, jogadasLivres map[string]int) bool {
	_, exists := jogadasLivres[jogadaEscolhida]
	if exists {
		delete(jogadasLivres, jogadaEscolhida)
		return true
	}
	return false
}
