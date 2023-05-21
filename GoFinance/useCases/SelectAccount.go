package useCases

import (
	"fmt"
	"strconv"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/utils"
)

func SelectAccount(accountsMap map[string]structs.Account) {
	bankKey := utils.GetBankTypeKey()
	account, ok := accountsMap[bankKey]
	if !ok {
		fmt.Println("Conta não encontrada")
		return
	}

	fmt.Println("Conta selecionada com com sucesso")
	ListAccount(account)

	for {
		utils.PrintMenuAccount()
		option := utils.GetInput("Digite a opção desejada:")
		utils.ClearConsole()
		switch option {
		case "1":
			fmt.Println("Saldo:", account.GetBalance())
		case "2":
			value, err := strconv.ParseFloat(utils.GetInput("Digite o valor do depósito"), 64)
			if err != nil {
				fmt.Println("Valor inválIdo")
				break
			}
			account.Deposit(value)
			accountsMap[bankKey] = account
		case "3":
			value, err := strconv.ParseFloat(utils.GetInput("Digite o valor do saque"), 64)
			if err != nil {
				fmt.Println("Valor inválIdo")
				break
			}
			message := account.Withdraw(value)
			fmt.Println(message, "Saldo:", account.GetBalance())
			accountsMap[bankKey] = account
		case "4":
			bankKeyDestiny := utils.GetBankTypeKey()
			destinyAccount, ok := accountsMap[bankKeyDestiny]
			if !ok {
				fmt.Println("Conta não encontrada")
				break
			}
			value, err := strconv.ParseFloat(utils.GetInput("Digite o valor da Transferência"), 64)
			if err != nil {
				fmt.Println("Valor inválIdo")
				break
			}
			account.Transfer(value, &destinyAccount)
			accountsMap[bankKey] = account
			accountsMap[bankKeyDestiny] = destinyAccount
		case "0":
			return
		default:
			fmt.Println("Opção inválIda")
		}
	}
}