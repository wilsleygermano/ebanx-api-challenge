package usecases

import (
	am "ebanx.api/account/data/models"
	ar "ebanx.api/account/domain/repositories"
)

func ChangeFunds(id string, ammount float64) map[string]any {
	repository := &ar.AccountRepository{}

	if id == "" {
		return map[string]any{
			"error":   true,
			"message": "account id is required",
		}
	}

	account := am.Account{
		ID:      id,
		Balance: ammount,
	}

	account = repository.AddAccount(account)

	return map[string]any{
		"ammount": account.Balance,
	}
}
