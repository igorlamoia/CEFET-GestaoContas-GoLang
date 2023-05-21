package useCases

import (
	"fmt"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/utils"
)

func RemoveAccount(accountsMap map[string]structs.Account) {
	document := utils.GetInput("Digite o documento da conta que deseja remover")
	_, ok := accountsMap[document]
	if !ok {
		fmt.Println("Conta não encontrada")
		return
	}
	delete(accountsMap, document)
	fmt.Println("Conta removIda com sucesso")
}