package config

import (
  "github.com/jinzhu/gorm"
  "github.com/gin-gonic/gin"

  "log"
  "net/http"

)


type Server struct {
  DB *gorm.DB
  Router *gin.Engine
}

func (server Server) Run(addr string) {
  log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) Initialize() {
  server.Router = gin.Default()
  server.InitializeRoutes()
}
