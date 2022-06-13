package main

func IniciaJogo() (bool, int) {
	// var rodadas = 9

	ExecutaJogada(jogador1.player, jogador1.caractere)
	ExecutaJogada(jogador2.player, jogador2.caractere)
	ExecutaJogada(jogador1.player, jogador1.caractere)
	ExecutaJogada(jogador2.player, jogador2.caractere)

	if ExecutaJogada(jogador1.player, jogador1.caractere) {
		return true, jogador1.player
	}

	if ExecutaJogada(jogador2.player, jogador2.caractere) {
		return true, jogador2.player
	}

	if ExecutaJogada(jogador1.player, jogador1.caractere) {
		return true, jogador1.player
	}

	if ExecutaJogada(jogador2.player, jogador2.caractere) {
		return true, jogador2.player
	}

	if ExecutaJogada(jogador1.player, jogador1.caractere) {
		return true, jogador1.player
	}
	return false, 0
}
