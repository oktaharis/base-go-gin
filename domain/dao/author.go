package dao

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;"`
	Fullname  string    `gorm:"size:56;not null;"`
	Gender    string    `gorm:"type:enum('m', 'f');not null;"`
	BirthDate time.Time `gorm:"type:datetime"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
