// D:\SEKOLAH\DIGI UP\GO\base-go-gin\domain\dao\book.go
package dao

import (
    "time"

    "gorm.io/gorm"
)

type Book struct {
    ID          uint           `gorm:"primaryKey;autoIncrement;"`
    Title       string         `gorm:"size:56;"`
    Subtitle    string         `gorm:"size:56;"`
    AuthorID    uint           `gorm:"not null;"`
    PublisherID uint           `gorm:"not null;"`
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}