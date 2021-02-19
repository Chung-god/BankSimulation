package main

import (
	"fmt"

	"github.com/Chung-god/bankSimulation/accounts"
)

func main() {
	account := accounts.NewAccounts("nico")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}

	fmt.Println(account.Balance())

	account.ChangeOwener("Dain")
	fmt.Println(account.Owener())
	fmt.Println(account)
}
