package dao

import (
	"base-gin/domain"
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	AccountID *uint
	Account   *Account          `gorm:"foreignKey:AccountID;"`
	Fullname  string            `gorm:"not null;"`
	Gender    domain.TypeGender `gorm:"type:enum('f','m');not null;"`
	BirthDate time.Time         `gorm:"not null;"`
}
