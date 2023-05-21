package useCases

import (
	"fmt"

	"gofinance.com/GoFinance/structs"
)

func ListAccounts(accountsMap map[string]structs.Account) {
	if len(accountsMap) == 0 {
		fmt.Println("Não há contas cadastradas")
		return
	}
	i := 1
	for key, account := range accountsMap {
		fmt.Println("----------------------")
		fmt.Println("Conta", i)
		fmt.Println("Banco", key)
		ListAccount(account)
		fmt.Println("----------------------")
		i++
	}
}

func ListAccount (account structs.Account) {
	fmt.Println("Id:", account.Id)
	fmt.Println("Nome:", account.GetNickName())
	fmt.Println("Tipo:", account.Owner.GetOwnerType())
	fmt.Println("Documento:", account.GetDocument())
	fmt.Println("Saldo:", account.GetBalance())
	fmt.Println("Chaves:", account.Keys)
}

func ListBanks (accountsMap map[string]structs.Account) {
	if len(accountsMap) == 0 {
		fmt.Println("Não há bancos cadastrados")
		return
	}
	fmt.Println("----------------------")
	fmt.Println("Bancos e contas cadastradas:")
	for key, _ := range accountsMap {
		fmt.Println("Banco", key)
	}
	fmt.Println("----------------------")
}