package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(message)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Falha ao ler dados:", err)
		os.Exit(1)
	}
	value = strings.TrimRight(value, "\r\n") // Remove newline characters from input
	return value
}

func PrintMenuManager() {
	fmt.Println(`
		1 - Criar conta\n
		2 - Entrar em conta\n
		3 - Listar contas\n
		4 - Remover conta\n
		0 - Sair
	`)
}

func PrintMenuAccount() {
	fmt.Println(`
		1 - Saldo\n
		2 - Depósito\n
		3 - Saque\n
		4 - Transferência\n
		0 - Voltar para gerenciamento de contas\n
	`)
}
