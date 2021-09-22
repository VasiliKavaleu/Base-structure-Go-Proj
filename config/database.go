package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
	"log"

	"goproj/models"
)

var DB *gorm.DB

type ConfigDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func (s *Server) ConnectDB(cfg *ConfigDB) {

	var err error
	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		fmt.Printf("Cannot connect to %s database ", cfg.DBName)
		log.Fatal("There is an error of connecting to postgres:", err)
	}

	fmt.Println("Connection to database is successful.")

	DB.AutoMigrate(models.Category{}, models.Item{})

}
