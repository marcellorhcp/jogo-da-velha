package main

func MarcaPosicaoEscolhida(posicaoEscolhida string, caracter string) {
	switch posicaoEscolhida {
	case "1":
		matriz[0][0] = caracter
	case "2":
		matriz[0][1] = caracter
	case "3":
		matriz[0][2] = caracter
	case "4":
		matriz[1][0] = caracter
	case "5":
		matriz[1][1] = caracter
	case "6":
		matriz[1][2] = caracter
	case "7":
		matriz[2][0] = caracter
	case "8":
		matriz[2][1] = caracter
	case "9":
		matriz[2][2] = caracter
	}
}
