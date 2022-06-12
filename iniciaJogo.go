package main

func IniciaJogo() bool {
	jogadaJogador1()
	jogadaJogador2()
	jogadaJogador1()
	jogadaJogador2()
	jogadaJogador1()

	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	jogadaJogador2()
	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	jogadaJogador1()
	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	jogadaJogador2()
	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	jogadaJogador1()
	if VerificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}
	return vencedor
}
