package dao

import (
	"time"

	"gorm.io/gorm"
)

type FormData struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;"`
	RadioButton string         `gorm:"size:56;not null;"`
	CheckBox    bool           `gorm:"not null;"`
	InputText   string         `gorm:"not null;"`
	SelectBox   string         `gorm:"not null;"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
