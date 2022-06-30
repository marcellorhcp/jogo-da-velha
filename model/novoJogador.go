package model

type Jogador struct {
	Numero    int
	Caractere string
}

func NovoJogador(player int, caractere string) Jogador {
	return Jogador{player, caractere}
}
