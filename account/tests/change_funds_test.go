package tests

import (
	repo "ebanx.api/account/domain/repositories"
	cf "ebanx.api/account/domain/use_cases"
	"testing"
	"sync"
)

func TestChangeFundsWithoutID(t *testing.T) {
	err := cf.ChangeFunds("", 10.0)
	if err == nil {
		t.Fatalf("Expected error, ID is empty")
	}
}

func TestCreateAnAccount(t *testing.T) {
	err := cf.ChangeFunds("123", 10.0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

var instance *repo.AccountRepository
var once sync.Once

func GetInstance() *repo.AccountRepository {
    once.Do(func() {
        instance = &repo.AccountRepository{}
    })
    return instance
}

func ResetAccounts() {
	repository := GetInstance()
	repository.ResetAccounts()
}

func TestIncreaseBalance(t *testing.T) {
	ResetAccounts()
	repository := GetInstance()
	err := cf.ChangeFunds("123", 10.0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	err2 := cf.ChangeFunds("123", 10.0)
	if err2 != nil {
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
