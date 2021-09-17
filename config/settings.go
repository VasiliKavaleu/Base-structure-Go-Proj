package config

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/gin-gonic/gin"

  "log"
  "net/http"
  "fmt"

  "goproj/models"

)


type ConfigDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Server struct {
  DB *gorm.DB
  Router *gin.Engine
}

func (server Server) Run(addr string) {
  fmt.Printf("Server running on: %s \n", addr)
  log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) Initialize(cfg *ConfigDB) {
  server.Router = gin.Default()
  server.InitializeRoutes()
  server.ConnectDB(cfg)
}

func (server *Server) ConnectDB(cfg *ConfigDB) {
  var err error
  server.DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
  cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
  if err != nil {
    // panic("Не удалось подключиться к базе данных")
    fmt.Printf("Cannot connect to %s database")
    log.Fatal("This is the error connecting to postgres:", err)
  }

  fmt.Println("Connection to database is successful.")
  server.DB.AutoMigrate(models.Category{}, models.Item{})

}
