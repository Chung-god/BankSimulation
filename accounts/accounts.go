package accounts

import (
	"errors"
	"fmt"
)

//There are no constructor so you need to make function(make struct)

//Account struct
type Account struct {
	owner   string
	balance int
}

//NewAccounts make account
func NewAccounts(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Balance show
func (a Account) Balance() int {
	return a.balance
}

//Deposit amount in account
func (a *Account) Deposit(amount int) {
	a.balance += amount

}

//Withdraw amount in account
func (a *Account) Withdraw(amount int) error {
	if amount > a.Balance() {
		return errors.New("You don't have a enough money")
	}
	a.balance -= amount
	return nil
}

//ChangeOwener of account
func (a *Account) ChangeOwener(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owener() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owener(), "'s account has ", a.Balance())
}
