package main

func VerificaVencedor(m [][]string) bool {
	switch {
	case m[0][0] == m[0][1] && m[0][1] == m[0][2] && m[0][2] != " ":
		return true
	case m[1][0] == m[1][1] && m[1][1] == m[1][2] && m[1][2] != " ":
		return true
	case m[2][0] == m[2][1] && m[2][1] == m[2][2] && m[2][2] != " ":
		return true
	case m[0][0] == m[1][0] && m[1][0] == m[2][0] && m[2][0] != " ":
		return true
	case m[0][1] == m[1][1] && m[1][1] == m[2][1] && m[2][1] != " ":
		return true
	case m[0][2] == m[1][2] && m[1][2] == m[2][2] && m[2][2] != " ":
		return true
	case m[0][0] == m[1][1] && m[1][1] == m[2][2] && m[2][2] != " ":
		return true
	case m[2][0] == m[1][1] && m[1][1] == m[0][2] && m[0][2] != " ":
		return true
	default:
		return false
	}
}
