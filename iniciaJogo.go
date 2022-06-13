package main

func IniciaJogo() (bool, int) {
	// var rodadas = 9

	ProximaJogada(jogador1.player, jogador1.caractere)
	ProximaJogada(jogador2.player, jogador2.caractere)
	ProximaJogada(jogador1.player, jogador1.caractere)
	ProximaJogada(jogador2.player, jogador2.caractere)

	ProximaJogada(jogador1.player, jogador1.caractere)
	if VerificaVencedor(matriz) {
		return true, jogador1.player
	}

	ProximaJogada(jogador2.player, jogador2.caractere)
	if VerificaVencedor(matriz) {
		return true, jogador2.player
	}

	ProximaJogada(jogador1.player, jogador1.caractere)
	if VerificaVencedor(matriz) {
		return true, jogador1.player
	}

	ProximaJogada(jogador2.player, jogador2.caractere)
	if VerificaVencedor(matriz) {
		return true, jogador2.player
	}

	ProximaJogada(jogador1.player, jogador1.caractere)
	if VerificaVencedor(matriz) {
		return true, jogador1.player
	}
	return false, 0
}
