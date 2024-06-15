package usecases

import (
	ar "ebanx.api/account/domain/repositories"
)

func ChangeFunds(id string, ammount float64) map[string]any {
	if id == "" {
		return map[string]any{
			"error":   true,
			"message": "account id is required",
		}
	}

	repository := &ar.AccountRepository{}
	// get account by id
	account := repository.GetAccountByID(id)

	if account.ID == "" {
		account.ID = id
		account.Balance = ammount
	} else {
		account.Balance += ammount
	}

	account = repository.AddAccount(account)

	return map[string]any{
		"ammount": account.Balance,
	}
}
