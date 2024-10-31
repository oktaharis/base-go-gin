package repository

import "base-gin/storage"

var (
	accountRepo *AccountRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = NewAccountRepository(db)
}

func GetAccountRepo() *AccountRepository {
	return accountRepo
}
