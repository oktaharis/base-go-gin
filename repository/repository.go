package repository

import "base-gin/storage"

var (
	accountRepo   *AccountRepository
	personRepo    *PersonRepository
	publisherRepo *PublisherRepository
	authorRepo    *AuthorRepository
	bookRepo      *BookRepository
	formDataRepo  *FormDataRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = NewAccountRepository(db)
	personRepo = NewPersonRepository(db)
	publisherRepo = NewPublisherRepository(db)
	authorRepo = NewAuthorRepository(db)
	bookRepo = NewBookRepository(db)
	formDataRepo = NewFormDataRepository(db)
}

func GetAccountRepo() *AccountRepository {
	return accountRepo
}

func GetPersonRepo() *PersonRepository {
	return personRepo
}

func GetPublisherRepo() *PublisherRepository {
	return publisherRepo
}

func GetAuthorRepo() *AuthorRepository {
	return authorRepo
}
func GetBookRepo() *BookRepository {
	return bookRepo
}
func GetFormDataRepo() *FormDataRepository {
	return formDataRepo
}
