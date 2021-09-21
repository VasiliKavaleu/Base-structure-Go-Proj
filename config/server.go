package config

import (
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/gin-gonic/gin"

  "log"
  "net/http"
  "fmt"

)


type Server struct {}


func (s *Server) Run(addr string, router *gin.Engine) {
  fmt.Printf("Server running on: %s \n", addr)
  log.Fatal(http.ListenAndServe(addr, router))
}
