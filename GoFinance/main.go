package main

import (
	"fmt"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/useCases"
	"gofinance.com/GoFinance/utils"
)

func main() {
	fmt.Println("Seja bem vindo de volta ao GoFinance!")
	manager()
}

func manager() {
	accountsMap := make(map[string]structs.Account)
	for {
		utils.PrintMenuManager()
		option := utils.GetInput("Digite a opção desejada:")
		utils.ClearConsole()
		switch option {
		case "1":
			useCases.AddAccount(accountsMap)
		case "2":
			useCases.SelectAccountAndMakeOperations(accountsMap)
		case "3":
			useCases.ListAccounts(accountsMap)
		case "4":
			useCases.RemoveAccount(accountsMap)
		case "0":
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}
