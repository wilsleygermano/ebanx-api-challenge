package tests

import (
	cf "ebanx.api/account/domain/use_cases"
	"testing"
	tools "ebanx.api/account/tests/tools"
)

func TestChangeFundsWithoutID(t *testing.T) {
	err := cf.ChangeFunds("", 10.0)
	if err == nil {
		t.Fatalf("Expected error, ID is empty")
	}
}

func TestCreateAnAccount(t *testing.T) {
	err := cf.ChangeFunds("123", 10.0)
	if err["error"] != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func ResetAccounts() {
	repository := tools.GetInstance()
	repository.ResetAccounts()
}

func TestIncreaseBalance(t *testing.T) {
	ResetAccounts()
	repository := tools.GetInstance()
	err := cf.ChangeFunds("123", 10.0)
	if err["error"] != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	err2 := cf.ChangeFunds("123", 10.0)
	if err2["error"] != nil {
		t.Fatalf("Expected no error, got %v", err2)
	}
	accounts := repository.GetAllAccounts()
	for _, account := range accounts {
		if account.ID == "123" {
			if account.Balance != 20.0 {
				t.Fatalf("Expected balance to be 20.0, got %v", account.Balance)
			}
		}
	}
}
