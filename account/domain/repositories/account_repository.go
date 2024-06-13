package repositories

import (
	ads "ebanx.api/account/data/data_sources"
	am "ebanx.api/account/data/models"
)

type AccountRepository struct {
	AccountsDataSource *ads.AccountsDataSource
}

func (ar *AccountRepository) GetAllAccounts() []am.Account {
	datasource := ar.GetDataSourceInstance()
	return datasource.GetAllAccounts()
}

func (ar *AccountRepository) AddAccount(account am.Account) {
	datasource := ar.GetDataSourceInstance()
	datasource.WriteAccount(account)
}

func (ar *AccountRepository) ResetAccounts() {
	datasource := ar.GetDataSourceInstance()
	datasource.ResetAccounts()
}

func (ar *AccountRepository) GetDataSourceInstance() *ads.AccountsDataSource {
    return ads.GetInstance()
}