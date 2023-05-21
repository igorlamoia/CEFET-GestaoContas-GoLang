package structs

import "fmt"

type Account struct {
	Id      int
	Keys    [5]string
	Balance float64
	Owner   Owner
}

func (a Account) GetDocument() string {
	return a.Owner.GetDocument()
}
func (a Account) GetNickName() string {
	return a.Owner.GetNickName()
}
func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) Withdraw(value float64) bool {
	if a.Balance < value {
		fmt.Println("Saldo insuficiente")
		return false
	}
	a.Balance -= value
	fmt.Println("Saque realizado com sucesso")
	fmt.Println("Saldo atual:", a.Balance)
	return true
}

func (a *Account) Deposit(value float64) bool {
	if value > 0 {
		a.Balance += value
		fmt.Println("Depósito realizado com sucesso", a.Balance)
		return true
	}
	fmt.Println("Valor inválido", a.Balance)
	return false
}

func (a *Account) Transfer(value float64, destiny *Account) bool {
	if value < a.Balance && value > 0 {
		a.Balance -= value
		destiny.Balance += value
		fmt.Println("Transferência realizada com sucesso")
		return true
	}
	fmt.Println("Saldo insuficiente")
	return false
}
