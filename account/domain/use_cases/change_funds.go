package usecases

import (
	am "ebanx.api/account/data/models"
	ar "ebanx.api/account/domain/repositories"
	"errors"
)

func ChangeFunds(id string, ammount float64) error {
	repository := &ar.AccountRepository{}

	if id == "" {
		return errors.New("you must provide a valid ID")
	}

	account := am.Account{
		ID:      id,
		Balance: ammount,
	}
	repository.AddAccount(account)

	return nil
}
