package usecases

import (
	r "ebanx.api/account/domain/repositories"
)

func TransferFunds(origin string, destination string, ammount float64) map[string]any {
	repository := &r.AccountRepository{}
	originAccount := repository.GetAccountByID(origin)
	destinationAccount := repository.GetAccountByID(destination)
	if originAccount.ID == "" {
		return map[string]any{
			"error": true,
		}
	}
	if destinationAccount.ID == "" {
		destinationAccount.ID = destination
		destinationAccount.Balance = 0
	}
	if originAccount.Balance < ammount {
		return map[string]any{
			"error": true,
		}
	}
	originAccount.Balance -= ammount
	destinationAccount.Balance += ammount
	repository.AddAccount(originAccount)
	repository.AddAccount(destinationAccount)

	return map[string]any{
		"origin": map[string]any{
			"id":      originAccount.ID,
			"balance": originAccount.Balance,
		},
		"destination": map[string]any{
			"id":      destinationAccount.ID,
			"balance": destinationAccount.Balance,
		},
	}
}
