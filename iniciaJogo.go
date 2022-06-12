package main

func IniciaJogo() bool {
	ProximaJogada(1)
	ProximaJogada(2)
	ProximaJogada(1)
	ProximaJogada(2)
	ProximaJogada(1)

	if VerificaVencedor(matriz) {
		return true
	}

	ProximaJogada(2)
	if VerificaVencedor(matriz) {
		return true
	}

	ProximaJogada(1)
	if VerificaVencedor(matriz) {
		return true
	}

	ProximaJogada(2)
	if VerificaVencedor(matriz) {
		return true
	}

	ProximaJogada(1)
	if VerificaVencedor(matriz) {
		return true
	}
	return false
}
