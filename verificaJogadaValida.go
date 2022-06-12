package main

func VerificaJogadaValida(j string) bool {
	_, exists := jogadasLivres[j]
	if exists {
		delete(jogadasLivres, j)
		return true
	}
	return false
}
