package useCases

import (
	"fmt"
	"strconv"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/utils"
)

func Deposit(accountsMap map[string]structs.Account, bankKey string, account structs.Account) {
	if len(accountsMap) == 0 {
		fmt.Println("Não há contas cadastrados")
		return
	}
	value, err := strconv.ParseFloat(utils.GetInput("Digite o valor do depósito"), 64)
	if err != nil {
		fmt.Println("Valor inválIdo")
		return
	}
	account.Deposit(value)
	accountsMap[bankKey] = account
}