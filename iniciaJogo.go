package main

func IniciaJogo() bool {
	ProximaJogada(1)
	ProximaJogada(2)
	ProximaJogada(1)
	ProximaJogada(2)
	ProximaJogada(1)

	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	ProximaJogada(2)
	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	ProximaJogada(1)
	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	ProximaJogada(2)
	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	ProximaJogada(1)
	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}
	return vencedor
}
