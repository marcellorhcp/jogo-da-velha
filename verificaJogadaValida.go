package main

func VerificaJogadaValida(jogadaEscolhida string, jogadasLivres map[string]int) bool {
	_, exists := jogadasLivres[jogadaEscolhida]
	if exists {
		delete(jogadasLivres, jogadaEscolhida)
		return true
	}
	return false
}
