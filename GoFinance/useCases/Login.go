package useCases

import (
	"fmt"
	"strconv"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/utils"
)

func Login(accountsMap map[string]structs.Account) {
	document := utils.GetInput("Digite o documento da conta")
	password := utils.GetInput("Digite a senha da conta")
	account, ok := accountsMap[document]
	if !ok {
		fmt.Println("Conta não encontrada")
		return
	}
	if account.Keys[0] != password {
		fmt.Println("Senha incorreta")
		return
	}
	fmt.Println("Login realizado com sucesso")
	for {
		utils.PrintMenuAccount()
		option := utils.GetInput("Digite a opção desejada:")
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
		case "3":
			value, err := strconv.ParseFloat(utils.GetInput("Digite o valor do saque"), 64)
			if err != nil {
				fmt.Println("Valor inválIdo")
				break
			}
			message := account.Withdraw(value)
			fmt.Println(message, "Saldo:", account.GetBalance())
		case "4":
			document := utils.GetInput("Digite o documento da conta de destino")
			destinyAccount, ok := accountsMap[document]
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
		case "0":
			return
		default:
			fmt.Println("Opção inválIda")
		}
	}
}