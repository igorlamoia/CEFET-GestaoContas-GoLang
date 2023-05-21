package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ( READ = '\n' )

// type Role string
// const (
// 	Unknown   Role = ""
// 	Guest     Role = "guest"
// 	Member    Role = "member"
// 	Moderator Role = "moderator"
// 	Admin     Role = "admin"
// )

type Owner interface {
	getDocument() string
	getNickname() string
}
type Person struct {
	name string
	cpf string
}
func (p Person) getDocument() string {
	return p.cpf
}
func (p Person) getNickname() string {
	return p.name
}
type Company struct {
	fantasyName string
	cpnpj string
}

func (c Company) getDocument() string {
	return c.cpnpj
}
func (c Company) getNickname() string {
	return c.fantasyName
}


type Account struct {
	id int
	keys [5]string
	balance float64
	owner Owner
}
func (a Account) getDocument() string {
	return a.owner.getDocument()
}
func (a Account) getNickname() string {
	return a.owner.getNickname()
}

func (a *Account) withdraw(value float64) string {
	if a.balance < value {
		return "Saldo insuficiente"
	}
	a.balance -= value
	return "Saque realizado com sucesso"
}

func (a *Account) deposit(value float64) (string, float64) {
	if value > 0 {
		a.balance += value
		return "Depósito realizado com sucesso", a.balance
	}
	return "Valor inválido", a.balance
}

func (a *Account) transfer(value float64, destiny *Account) {
	if value < a.balance && value > 0 {
		a.balance -= value
		destiny.balance += value
		fmt.Println("Transferência realizada com sucesso")
		return
	}
	fmt.Println("Saldo insuficiente")
}

func (a *Account) getBalance() float64 {
	return a.balance
}

func printMenu() {
	fmt.Println("1 - Saldo\n2 - Depósito\n3 - Saque\n4 - Transferência\n0 - Sair")
}

func main() {
  fmt.Println("Seja bem vindo de volta ao MyWallet!")
	acc := Account{1,[5]string{"igorlamoia@hotmail.com"}, 100.0, Person{"John", "123.456.789-00"}}
	acc2 := Account{2,[5]string{"igorlamoia@hotmail.com"}, 100.0, Company{"Company", "123.456.789/0001-00"}}
	fmt.Println(acc.getDocument())
	fmt.Println(acc.getNickname())
	fmt.Println(acc2.getDocument())
	fmt.Println(acc2.getNickname())
	printMenu()
	fmt.Println(acc.withdraw(10.0))
	fmt.Println(acc.getBalance())
	fmt.Println(acc.deposit(10.0))
	fmt.Println(acc.getBalance())
	acc.transfer(10.0, &acc2)
	fmt.Println(acc.getBalance())
	fmt.Println(acc2.getBalance())
	acc.transfer(100.0, &acc2)
	acc.transfer(89.0, &acc2)
	fmt.Println(acc.getBalance())
	fmt.Println(acc2.getBalance())

	reader := bufio.NewReader(os.Stdin)

	for {
		printMenu()

		optionStr, err := reader.ReadString(READ)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Remove newline characters from input
		optionStr = strings.TrimRight(optionStr, "\r\n")

		number, err := strconv.Atoi(optionStr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Invalid option. Please try again.")
			continue
		}

		switch number {
		case 1:
			fmt.Println("Showing balance...")
			// Implement the logic for showing balance
		case 2:
			fmt.Print("Enter the amount to deposit: ")
			amountStr, err := reader.ReadString(READ)
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			amountStr = amountStr[:len(amountStr)-1]
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Invalid amount. Please try again.")
				continue
			}
			fmt.Println("Depositing amount:", amount)
			// Implement the logic for depositing
		case 3:
			fmt.Print("Enter the amount to withdraw: ")
			amountStr, err := reader.ReadString(READ)
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			amountStr = amountStr[:len(amountStr)-1]
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Invalid amount. Please try again.")
				continue
			}
			fmt.Println("Withdrawing amount:", amount)
			// Implement the logic for withdrawing
		case 4:
			fmt.Print("Enter the amount to transfer: ")
			amountStr, err := reader.ReadString(READ)
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			amountStr = amountStr[:len(amountStr)-1]
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Invalid amount. Please try again.")
				continue
			}
			fmt.Println("Transferring amount:", amount)
			// Implement the logic for transferring
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

}
