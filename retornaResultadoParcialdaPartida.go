package main

import "fmt"

//RetornaResultadoParcialdaPartida possibilita ao jogador ter uma visão do andamento da partida antes de decidir a sua jogada
func RetornaResultadoParcialdaPartida(matriz [][]string) string {
	return fmt.Sprintf(
		`		+---+---+---+
		| %s | %s | %s |
		+---+---+---+
 		| %s | %s | %s |
		+---+---+---+
 		| %s | %s | %s |
		+---+---+---+`,
		matriz[2][0], matriz[2][1], matriz[2][2],
		matriz[1][0], matriz[1][1], matriz[1][2],
		matriz[0][0], matriz[0][1], matriz[0][2])
}
