package model

//SelecionaPosicao recebe uma jogada e retorna os valores correspondentes para a matriz do jogo da velha
func SelecionaPosicao(jogada string) (int, int) {
	switch jogada {
	case "1":
		return 0, 0
	case "2":
		return 0, 1
	case "3":
		return 0, 2
	case "4":
		return 1, 0
	case "5":
		return 1, 1
	case "6":
		return 1, 2
	case "7":
		return 2, 0
	case "8":
		return 2, 1
	case "9":
		return 2, 2
	default:
		return -1, -1
	}
}
