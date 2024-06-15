package tests

import (
	uc "ebanx.api/account/domain/use_cases"
	tools "ebanx.api/account/tests/tools"
	"testing"
)

func TestWithdrawFunds(t *testing.T) {
	repository := tools.GetInstance()
	err := uc.ChangeFunds("123", 10.0)
	if err["error"] != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	err2 := uc.WithdrawFunds("123", 5.0)
	if err2["error"] != nil {
		t.Fatalf("Expected no error, got %v", err2)
	}

	accounts := repository.GetAllAccounts()
	for _, account := range accounts {
		if account.ID == "123" {
			if account.Balance != 5.0 {
				t.Fatalf("Expected balance to be 5.0, got %v", account.Balance)
			}
		}
	}
}
