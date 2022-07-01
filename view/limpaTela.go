package view

import (
	"fmt"
	"os"
	"os/exec"
)

func LimpaTela() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func EnterParaContinuar() {
	var tecla string
	fmt.Print("\nPressione Enter para continuar")
	fmt.Scanf("%s", &tecla)
	LimpaTela()
}
