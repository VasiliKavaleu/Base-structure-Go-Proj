package models


import (
	// "errors"
	// "fmt"
	"time"

	// "github.com/jinzhu/gorm"
)

type Category struct {
  ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
  Title       string    `gorm:"size:255;not null;unique" json:"title"`
  CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
  UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
