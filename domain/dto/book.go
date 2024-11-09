// D:\SEKOLAH\DIGI UP\GO\base-go-gin\domain\dto\book.go
package dto

import "base-gin/domain/dao"

type BookCreateReq struct {
	Title       string `json:"title" binding:"required,min=2,max=48"`
	Subtitle    string `json:"subtitle" binding:"min=2,max=48"`
	AuthorID    uint   `json:"author_id" binding:"required"`
	PublisherID uint   `json:"publisher_id" binding:"required"`
}

func (o *BookCreateReq) ToEntity() dao.Book {
	var item dao.Book
	item.Title = o.Title
	item.Subtitle = o.Subtitle
	item.AuthorID = o.AuthorID
	item.PublisherID = o.PublisherID

	return item
}

type BookResp struct {
	ID          uint    `json:"id"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`     // Perbaiki typo "Subtilte"
	AuthorID    uint   `json:"author_id"`    // Gunakan snake_case untuk konsistensi
	PublisherID uint   `json:"publisher_id"` // Gunakan snake_case untuk konsistensi
}

func (o *BookResp) FromEntity(item *dao.Book) {
	o.ID = item.ID
	o.Title = item.Title
	o.Subtitle = item.Subtitle       // Tambahkan mapping untuk Subtitle
	o.AuthorID = item.AuthorID       // Tambahkan mapping untuk AuthorID
	o.PublisherID = item.PublisherID // Tambahkan mapping untuk PublisherID
}

type BookUpdateReq struct {
	ID          uint   `json:"-"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`     // Perbaiki typo "Subtilte"
	AuthorID    uint   `json:"author_id"`    // Gunakan snake_case untuk konsistensi
	PublisherID uint   `json:"publisher_id"` // Gunakan snake_case untuk konsistensi
}
