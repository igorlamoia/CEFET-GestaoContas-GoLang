package useCases

import (
	"fmt"
	"strconv"
	"strings"

	"gofinance.com/GoFinance/structs"
	"gofinance.com/GoFinance/utils"
)

func AddAccount(accountsMap map[string]structs.Account) {
	OwnerType := utils.GetInput("Digite o tipo de conta que deseja criar: PF (Pesso física) ou PJ (Pessoa Jurídica?")
	OwnerType = strings.ToUpper(OwnerType)
	if OwnerType != "PF" && OwnerType != "PJ" {
		fmt.Println("Tipo de conta inválida")
		return
	}
	// celular := utils.GetInput("Digite a senha da conta")

	if(OwnerType == "PF"){
		Name := utils.GetInput("Digite o nome do proprietário da conta")
		Email := utils.GetInput("Digite o email principal da conta")
		Document := utils.GetInput("Digite o CPF da conta")
		Bank := utils.GetInput("Digite o banco da conta")
		Balance, err := strconv.ParseFloat(utils.GetInput("Valor do depósito inicial da conta"), 64)
		if err != nil {
			fmt.Println("Valor inválIdo")
			return
		}
		Id := utils.GetRandomNumber()
		accountsMap[Bank+"_PF"] = structs.Account{Id, [5]string{Email}, Balance, structs.Person{Name, Document}}
		return
	}
	Name := utils.GetInput("Digite o nome fantasia da empresa")
	Email := utils.GetInput("Digite o email principal da conta")
	Document := utils.GetInput("Digite o CNPJ da empresa")
	Bank := utils.GetInput("Digite o banco da conta da empresa")
	Balance, err := strconv.ParseFloat(utils.GetInput("Valor do depósito inicial da conta"), 64)
	if err != nil {
		fmt.Println("Valor inválIdo")
		return
	}
	Id := utils.GetRandomNumber()
	accountsMap[Bank+"_PJ"] = structs.Account{Id, [5]string{Email}, Balance, structs.Company{Name, Document}}
}