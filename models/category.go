package models

import (
	// "errors"
	// "fmt"
	"time"

	"github.com/jinzhu/gorm"
	//"github.com/gin-gonic/gin"
)

type Category struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Category) FindAllCategories(db *gorm.DB) (*[]Category, error) {
	var err error
	categories := []Category{}
	err = db.Debug().Model(&Category{}).Limit(20).Order("created_at desc").Find(&categories).Error
	if err != nil {
		return &[]Category{}, err
	}
	return &categories, nil
}

func (c *Category) SaveCategory(db *gorm.DB) (*Category, error) {
	var err error
	err = db.Debug().Model(&Category{}).Create(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil

}

func (c *Category) FindCategoryById(db *gorm.DB, cid uint64) (*Category, error) {
	err := db.Debug().Model(&Category{}).Where("id = ?", cid).Take(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c *Category) DeleteACategory(db *gorm.DB) (int64, error) {
	db = db.Debug().Model(&Category{}).Where("id = ?", c.ID).Take(&c).Delete(&Category{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (c *Category) UpdateACategory(db *gorm.DB) (*Category, error) {
	err := db.Debug().Model(&Category{}).Where("id = ?", c.ID).Updates(Category{Title: c.Title, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Category{}, err
	}

	return c, nil

}

