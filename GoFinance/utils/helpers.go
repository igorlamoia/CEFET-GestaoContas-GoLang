package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetBankTypeKey() string {
	bank := GetInput("Digite o nome do banco")
	ownerType := GetInput("Digite o tipo de conta: PF (Pessoa física) ou PJ (Pessoa Jurídica)")
	ownerType = strings.ToUpper(ownerType)
	if(ownerType != "PF" && ownerType != "PJ") {
		fmt.Println("Tipo de conta inválido")
		os.Exit(1)
	}
	return bank + "_" + strings.ToUpper(ownerType)
}

func GetRandomNumber() int {
	seed := time.Now().UnixNano() // Use current time as the seed
	randSrc := rand.NewSource(seed)
	randGen := rand.New(randSrc)
	min := 1000
	max := 9999

	return randGen.Intn(max-min+1) + min
}

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
