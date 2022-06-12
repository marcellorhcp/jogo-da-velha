package main

func MarcaPosicaoEscolhida(x string, n int) {
	var opcao string = "X"
	if n == 2 {
		opcao = "0"
	}
	switch x {
	case "1":
		matriz[0][0] = opcao
	case "2":
		matriz[0][1] = opcao
	case "3":
		matriz[0][2] = opcao
	case "4":
		matriz[1][0] = opcao
	case "5":
		matriz[1][1] = opcao
	case "6":
		matriz[1][2] = opcao
	case "7":
		matriz[2][0] = opcao
	case "8":
		matriz[2][1] = opcao
	case "9":
		matriz[2][2] = opcao
	}
}
