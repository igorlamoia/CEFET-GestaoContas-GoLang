package useCases

import (
	"fmt"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/utils"
)

func SelectAccountAndMakeOperations(accountsMap map[string]structs.Account) {
	if len(accountsMap) == 0 {
		fmt.Println("Não há contas cadastrados")
		return
	}
	ListBanks(accountsMap)
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
			Deposit(accountsMap, bankKey, account)
			return
		case "2":
			Withdraw(accountsMap, bankKey, account)
			return
		case "3":
			ListBanks(accountsMap)
			fmt.Println("Digite o banco de destino")
			Transfer(accountsMap, bankKey, account)
			return
		case "0":
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}