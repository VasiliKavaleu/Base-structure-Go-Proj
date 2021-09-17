package models


import (
	// "errors"
	// "fmt"
	"time"

	// "github.com/jinzhu/gorm"
)

type Item struct {
  ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
  Title       string    `gorm:"size:255;not null;unique" json:"title"`
  Guarantee   uint8     `gorm: json:"guarantee"`
  Category    Category  `json:"category"`
  CategoryID  uint64    `gorm:"not null" json:"author_id"`
  CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
  UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
