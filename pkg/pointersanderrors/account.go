package pointersanderrors

import "errors"

type Dollar float64

type Account struct {
	Id      int
	Balance Dollar
}

type BankOperations interface {
	Transfer(account *Account, amount Dollar) error
}

func (a *Account) transfer(account *Account, amount Dollar) error {
	if amount > a.Balance {
		return errors.New("insufficient funds")
	}

	a.Balance -= amount
	account.Balance += amount

	return nil
}
