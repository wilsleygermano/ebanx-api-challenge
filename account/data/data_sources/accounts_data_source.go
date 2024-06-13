package datasources

import (
	am "ebanx.api/account/data/models"
	"sync"
)


type AccountsDataSource struct {
	accounts []am.Account
	mu       sync.Mutex
}

var instance *AccountsDataSource
var once sync.Once

func  GetInstance() *AccountsDataSource {
	once.Do(func() {
		instance = &AccountsDataSource{
			accounts: make([]am.Account, 0),
		}
	})
	return instance
}

func (ds *AccountsDataSource) GetAllAccounts() []am.Account {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	return append([]am.Account{}, ds.accounts...)
}

func (ds *AccountsDataSource) WriteAccount(account am.Account) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	// first we check if the account already exists
	for i, acc := range ds.accounts {
		if acc.ID == account.ID {
			ds.accounts[i].Balance += account.Balance
			return
		}
	}
	ds.accounts = append(ds.accounts, account)
}

func (ds *AccountsDataSource) ResetAccounts() {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.accounts = make([]am.Account, 0)
}
