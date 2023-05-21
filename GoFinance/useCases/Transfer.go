package useCases

import (
	"fmt"
	"strconv"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/utils"
)

func Transfer(accountsMap map[string]structs.Account, bankKey string, account structs.Account) {
	if len(accountsMap) == 0 {
		fmt.Println("Não há contas cadastrados")
		return
	}
	bankKeyDestiny := utils.GetBankTypeKey()
	destinyAccount, ok := accountsMap[bankKeyDestiny]
	if !ok {
		fmt.Println("Conta não encontrada")
		return
	}
	value, err := strconv.ParseFloat(utils.GetInput("Digite o valor da Transferência"), 64)
	if err != nil {
		fmt.Println("Valor inválido")
		return
	}
	account.Transfer(value, &destinyAccount)
	accountsMap[bankKey] = account
	accountsMap[bankKeyDestiny] = destinyAccount
}