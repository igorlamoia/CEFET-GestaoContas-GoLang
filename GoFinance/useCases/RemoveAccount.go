package useCases

import (
	"fmt"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/utils"
)

func RemoveAccount(accountsMap map[string]structs.Account) {
	if len(accountsMap) == 0 {
		fmt.Println("Não há contas cadastradas")
		return
	}
	fmt.Println("Excluindo conta")
	ListBanks(accountsMap)
	accountKey := utils.GetBankTypeKey()
	_, ok := accountsMap[accountKey]
	if !ok {
		fmt.Println("Conta não encontrada")
		return
	}
	delete(accountsMap, accountKey)
	fmt.Println("Conta removida com sucesso")
}