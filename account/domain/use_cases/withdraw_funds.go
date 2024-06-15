package usecases

import (
	r "ebanx.api/account/domain/repositories"
)

func WithdrawFunds(id string, ammount float64) map[string]any {
	repository := &r.AccountRepository{}
	// get account by id
	account := repository.GetAccountByID(id)
	if account.ID == "" {
		return map[string]any{
			"error":   true,
			"ammount": 0,
		}
	}
	if account.Balance < ammount {
		return map[string]any{
			"error":   true,
			"ammount": account.Balance,
		}
	}
	account.Balance -= ammount
	repository.AddAccount(account)
	return map[string]any{
		"ammount": account.Balance,
	}
}
