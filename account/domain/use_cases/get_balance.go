package usecases

import (
	r "ebanx.api/account/domain/repositories"
)

func GetBalance(id string) map[string]any {
	repository := &r.AccountRepository{}
	// get account by id
	account := repository.GetAccountByID(id)
	if account.ID == "" {
		return map[string]any{
			"error":   true,
			"ammount": 0,
		}
	}
	return map[string]any{
		"ammount": account.Balance,
	}
}
