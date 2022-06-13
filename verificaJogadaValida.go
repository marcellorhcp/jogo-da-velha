package main

func VerificaJogadaValida(jogadaEscolhida string) bool {
	_, exists := jogadasLivres[jogadaEscolhida]
	if exists {
		delete(jogadasLivres, jogadaEscolhida)
		return true
	}
	return false
}
