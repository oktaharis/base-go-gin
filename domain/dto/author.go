package dto

import (
	"base-gin/domain/dao"
	"time"
)

// req data baru
type AuthorCreateReq struct {
	Fullname  string    `json:"fullname" binding:"required,min=2,max=56"`
	Gender    string    `json:"gender" binding:"required,oneof=m f"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
}

func (a *AuthorCreateReq) ToEntity() dao.Author {
	var author dao.Author
	author.Fullname = a.Fullname
	author.Gender = a.Gender
	author.BirthDate = a.BirthDate

	return author
}

// response
type AuthorResp struct {
	ID        uint      `json:"id"`
	Fullname  string    `json:"fullname"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
}

func (a *AuthorResp) FromEntity(author *dao.Author) {
	a.ID = author.ID
	a.Fullname = author.Fullname
	a.Gender = author.Gender
	a.BirthDate = author.BirthDate
}

type AuthorUpdateReq struct {
	ID        uint      `json:"-"`
	Fullname  string    `json:"fullname" binding:"required,min=2,max=56"`
	Gender    string    `json:"gender" binding:"required,oneof=m f"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
}
