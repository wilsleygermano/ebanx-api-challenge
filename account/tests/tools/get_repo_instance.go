package tools

import (
	repo "ebanx.api/account/domain/repositories"
	"sync"
)

var instance *repo.AccountRepository
var once sync.Once

func GetInstance() *repo.AccountRepository {
	once.Do(func() {
		instance = &repo.AccountRepository{}
	})
	return instance
}
