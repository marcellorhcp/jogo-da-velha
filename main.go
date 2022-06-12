package main

import "fmt"

func retornaResultado(s string) string {
	return fmt.Sprintf(
		`		| %s | %s | %s |
		-------------
 		| %s | %s | %s |
		-------------
 		| %s | %s | %s |`, matriz[2][0], matriz[2][1], matriz[2][2], matriz[1][0], matriz[1][1], matriz[1][2], matriz[0][0], matriz[0][1], matriz[0][2])
}
func marcaPosicaoEscolhida(x string, n int) {
	var opcao string
	if n == 1 {
		opcao = "X"
	} else {
		opcao = "O"
	}
	switch x {
	case "1":
		matriz[0][0] = opcao
		fmt.Println("\n", retornaResultado(matriz[0][0]))
	case "2":
		matriz[0][1] = opcao
		fmt.Println("\n", retornaResultado(matriz[0][1]))
	case "3":
		matriz[0][2] = opcao
		fmt.Println("\n", retornaResultado(matriz[0][2]))
	case "4":
		matriz[1][0] = opcao
		fmt.Println("\n", retornaResultado(matriz[1][0]))
	case "5":
		matriz[1][1] = opcao
		fmt.Println("\n", retornaResultado(matriz[1][1]))
	case "6":
		matriz[1][2] = opcao
		fmt.Println("\n", retornaResultado(matriz[1][2]))
	case "7":
		matriz[2][0] = opcao
		fmt.Println("\n", retornaResultado(matriz[2][0]))
	case "8":
		matriz[2][1] = opcao
		fmt.Println("\n", retornaResultado(matriz[2][1]))
	case "9":
		matriz[2][2] = opcao
		fmt.Println("\n", retornaResultado(matriz[2][2]))
	}
}

func verificaVencedor(m [][]string) bool {
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
func verificaJogadaValida(j string) bool {
	_, exists := jogadasLivres[j]
	if exists {
		delete(jogadasLivres, j)
		return true
	}
	return false
}
func jogadaJogador1() {
	jogador = 1
	fmt.Println("\nJogador 1 - Faça uma jogada - X")
	fmt.Scan(&jogada)
	if verificaJogadaValida(jogada) {
		marcaPosicaoEscolhida(jogada, jogador)
	} else {
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		jogadaJogador1()
	}
}
func jogadaJogador2() {
	jogador = 2
	fmt.Println("\nJogador 2 - Faça uma jogada - O")
	fmt.Scan(&jogada)
	if verificaJogadaValida(jogada) {
		marcaPosicaoEscolhida(jogada, jogador)
	} else {
		fmt.Println("\nOpa, parece que você escolheu uma posição já ocupada")
		fmt.Println("\nou digitou um número / caracter inválido")
		jogadaJogador2()
	}
}
func iniciaJogo() bool {
	jogadaJogador1()
	jogadaJogador2()
	jogadaJogador1()
	jogadaJogador2()
	jogadaJogador1()

	if verificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	jogadaJogador2()
	if verificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	jogadaJogador1()
	if verificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	jogadaJogador2()
	if verificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}

	jogadaJogador1()
	if verificaVencedor(matriz) {
		vencedor = true
		return vencedor
	}
	return vencedor
}

var jogador int
var andamentoDoJogo string
var jogada string
var vencedor bool = false
var matriz = [][]string{
	{" ", " ", " "},
	{" ", " ", " "},
	{" ", " ", " "},
}
var jogadasLivres = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func main() {
	// var nome string
	// fmt.Println("Digite o nome do Jogador 1")
	// fmt.Scan(&nome)

	fmt.Printf(`
	# Bem-vindo ao Jogo da Velha
	# O jogo funciona da seguinte forma:
	# Cada número escolhido representa uma casa no jogo da velha
	# A ordem das casas está de acordo com o teclado númerico
	# Boa sorte
	
		| 7 | 8 | 9 |
		------------
		| 4 | 5 | 6 |
		------------
		| 1 | 2 | 3 |

`)
	iniciaJogo()
	if vencedor {
		fmt.Println("\nParabéns, você venceu!!! :)")
	} else {
		fmt.Println("\nih, deu velha :(")
	}
}
