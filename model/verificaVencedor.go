package model

//verificaHorizontal verifica se a combinação atual de jogadas na horizontal tem um vencedor
func verificaHorizontal(m [][]string) bool {
	switch {
	case m[0][0] == m[0][1] && m[0][1] == m[0][2] && m[0][2] != " ":
		return true
	case m[1][0] == m[1][1] && m[1][1] == m[1][2] && m[1][2] != " ":
		return true
	case m[2][0] == m[2][1] && m[2][1] == m[2][2] && m[2][2] != " ":
		return true
	default:
		return false
	}
}

//verificaVertical verifica se a combinação atual de jogadas na vertical tem um vencedor
func verificaVertical(m [][]string) bool {
	switch {
	case m[0][0] == m[1][0] && m[1][0] == m[2][0] && m[2][0] != " ":
		return true
	case m[0][1] == m[1][1] && m[1][1] == m[2][1] && m[2][1] != " ":
		return true
	case m[0][2] == m[1][2] && m[1][2] == m[2][2] && m[2][2] != " ":
		return true
	default:
		return false
	}
}

//verificaDiagonal verifica se a combinação atual de jogadas na diagonal tem um vencedor
func verificaDiagonal(m [][]string) bool {
	switch {
	case m[0][0] == m[1][1] && m[1][1] == m[2][2] && m[2][2] != " ":
		return true
	case m[2][0] == m[1][1] && m[1][1] == m[0][2] && m[0][2] != " ":
		return true
	default:
		return false
	}
}

//VerificaVencedor verifica se a combinação atual de jogadas marcadas na matriz tem um vencedor
func VerificaVencedor(m [][]string) bool {
	switch {
	case verificaHorizontal(m):
		return true
	case verificaVertical(m):
		return true
	case verificaDiagonal(m):
		return true
	default:
		return false
	}
}
