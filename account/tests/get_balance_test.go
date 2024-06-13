package tests

import (
	"testing"

	uc "ebanx.api/account/domain/use_cases"
)

func TestGetBalanceWithExistingAccount(t *testing.T) {
	ResetAccounts()
	errCF := uc.ChangeFunds("123", 10.0)
	if errCF != nil {
		t.Fatalf("Expected no error, got %v", errCF)
	}
	balance := uc.GetBalance("123")
	if balance["ammount"] != 10.0 {
		t.Fatalf("Expected balance to be 10.0, got %v", balance["ammount"])
	}

}

func TestGetBalanceWithoutAccount(t *testing.T) {
	balance := uc.GetBalance("9999")
	if balance["error"] == nil {
		t.Fatalf("Expected account not found error, got %v", balance["error"])
	}
}
