package repository

import (
	"base-gin/domain/dao"
	"base-gin/exception"
	"base-gin/storage"
	"errors"

	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Create(newItem *dao.Person) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Create(&newItem)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *PersonRepository) GetByAccountID(accountID uint) (dao.Person, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var item dao.Person
	tx := r.db.WithContext(ctx).Where(dao.Person{AccountID: &accountID}).
		First(&item)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return item, exception.ErrUserNotFound
		}

		return item, tx.Error
	}

	return item, nil
}
