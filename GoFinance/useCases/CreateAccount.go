package useCases

import (
	"fmt"
	"strconv"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/utils"
)

func CreateAccount(accountsMap map[string]structs.Account) {
	OwnerType := utils.GetInput("Digite o tipo de conta que deseja criar: PF ou PJ")
	if OwnerType != "PF" && OwnerType != "PJ" {
		fmt.Println("Tipo de conta inválIda")
		return
	}
	Name := utils.GetInput("Digite o nome do proprietário da conta")
	document := utils.GetInput("Digite o documento do proprietário da conta")
	Balance, err := strconv.ParseFloat(utils.GetInput("Digite o saldo inicial da conta"), 64)
	if err != nil {
		fmt.Println("Valor inválIdo")
		return
	}
	email := utils.GetInput("Digite o email da conta")
	password := utils.GetInput("Digite a senha da conta")

	accountsMap[document] = structs.Account{1, [5]string{email, password}, Balance, structs.Person{Name, document}}
}