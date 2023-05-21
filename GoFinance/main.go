package main

import (
	"fmt"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/useCases"
	"gofinance.com/GoFinance/utils"
)

const ( READ = '\n' )

func main() {
	fmt.Println("Seja bem vindo de volta ao GoFinance!")
	manager()
	return
	// accounts := make([]Account, 0, 5)
	// accounts = append(accounts, Account{1,[5]string{"igorlamoia@hotmail.com"}, 25000.0, Person{"Igor Lamoia", "123.456.789-00"}})
	// accounts = append(accounts, Account{2,[5]string{"hbifinanceira@hbi.com"}, 100.0, Company{"HBI - Financeira", "123.456.789/0001-00"}})
	// accounts = append(accounts, Account{3,[5]string{"milena@gmail.com"}, 9000.0, Person{"Milena", "123.456.789-22"}})

	// Gerenciar contas diferentes com o mesmo map
	accountsMap := make(map[string]structs.Account)
	accountsMap["bancoPAN-PF"] = structs.Account{1,[5]string{"igorlamoia@hotmail.com"}, 25000.0, structs.Person{"Igor Lamoia", "123.456.789-00"}}
	// for _, account := range accounts {
	// 	accountsMap[account.GetDocument()] = account
	// }
}

func listAccounts(accountsMap map[string]structs.Account) {
	for _, account := range accountsMap {
		fmt.Println(account)
	}
}

func manager() {
	accountsMap := make(map[string]structs.Account)
	for {
		utils.PrintMenuManager()
		option := utils.GetInput("Digite a opção desejada:")
		switch option {
		case "1":
			useCases.CreateAccount(accountsMap)
		case "2":
			useCases.Login(accountsMap)
		case "3":
			listAccounts(accountsMap)
		case "4":
			useCases.RemoveAccount(accountsMap)
		case "0":
			return
		default:
			fmt.Println("Opção inválIda")
		}
	}
}
