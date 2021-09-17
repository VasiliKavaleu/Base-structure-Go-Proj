package models


import (
	// "errors"
	// "fmt"
	"time"

	"github.com/jinzhu/gorm"
  //"github.com/gin-gonic/gin"
)

type Category struct {
  ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
  Title       string    `gorm:"size:255;not null;unique" json:"title"`
  CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
  UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (category *Category) FindAllCetegories(db *gorm.DB)(*[]Category, error){
  var err error
  categories := []Category{}
  err = db.Debug().Model(&Category{}).Limit(20).Order("created_at desc").Find(&categories).Error
  if err != nil {
    return &[]Category{}, err
  }
  return &categories, nil
}
