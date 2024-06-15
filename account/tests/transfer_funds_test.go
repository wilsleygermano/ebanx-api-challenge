package tests

import (
	uc "ebanx.api/account/domain/use_cases"
	tools "ebanx.api/account/tests/tools"
	"testing"
)

func TestTransferFunds(t *testing.T) {
	repository := tools.GetInstance()
	err := uc.ChangeFunds("100", 15.0)
	if err["error"] != nil {
		t.Fatalf("Expected no error, got %v", err)
	}


	err3 := uc.TransferFunds("100", "300", 15.0)
	if err3["error"] != nil {
		t.Fatalf("Expected no error, got %v", err3)
	}

	accounts := repository.GetAllAccounts()
	for _, account := range accounts {
		if account.ID == "100" {
			if account.Balance != 0.0 {
				t.Fatalf("Expected balance to be 0.0, got %v", account.Balance)
			}
		}
		if account.ID == "300" {
			if account.Balance != 15.0 {
				t.Fatalf("Expected balance to be 15.0, got %v", account.Balance)
			}
		}
	}

}
