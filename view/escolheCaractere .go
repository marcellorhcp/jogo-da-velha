package view

import (
	"fmt"

	"github.com/marcellorhcp/jogo-da-velha/model"
)

//EscolheCaractere recebe dois ponteiros do tipo jogador e possibilita a escolha do caractere que será utilizado
//a função é chamada de forma recursiva caso o jogador não escolha uma opção válida
func EscolheCaractere(jogador1, jogador2 *model.Jogador) {
	var escolha string
	fmt.Println("Jogador 1, escolha qual caracter será utilizado")
	fmt.Println("Digite 1 para jogar com X")
	fmt.Println("Digite 2 para jogar com O")
	fmt.Scan(&escolha)
	LimpaTela()

	switch escolha {
	case "1":
		jogador1.Caractere = "X"
		jogador2.Caractere = "O"
	case "2":
		jogador1.Caractere = "O"
		jogador2.Caractere = "X"
	default:
		LimpaTela()
		fmt.Println("Parece que você digitou uma opção inválida")
		EscolheCaractere(jogador1, jogador2)
	}
}
