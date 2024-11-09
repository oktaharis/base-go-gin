package repository

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/exception"
	"base-gin/storage"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(newBook *dao.Book) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Create(&newBook)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *BookRepository) GetById(id uint) (*dao.Book, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var book dao.Book
	tx := r.db.WithContext(ctx).First(&book, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, exception.ErrUserNotFound
		}
		return nil, tx.Error
	}

	return &book, nil
}

func (r *BookRepository) GetList(params dto.Filter) ([]dao.Book, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var books []dao.Book
	tx := r.db.WithContext(ctx)

	if params.Keyword != "" {
		q := fmt.Sprintf("%%%s%%", params.Keyword)
		tx = tx.Where("title LIKE ?", q)
	}
	if params.Start >= 0 {
		tx = tx.Offset(params.Start)
	}
	if params.Limit > 0 {
		tx = tx.Limit(params.Limit)
	}

	tx = tx.Order("title ASC").Find(&books)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}

	return books, nil
}

func (r *BookRepository) Update(params *dto.BookUpdateReq) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Model(&dao.Book{}).
		Where("id = ?", params.ID).
		Updates(map[string]interface{}{
			"Title":       params.Title,
			"Subtitle":    params.Subtitle,
			"author_id":    params.AuthorID,
			"publisher_id": params.PublisherID,
		})

	return tx.Error
}

func (r *BookRepository) Delete(id uint) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Delete(&dao.Book{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
