package controllers

import (
  "github.com/gin-gonic/gin"
  "goproj/models"
  "goproj/config/Server"
  "net/http"
)


func GetCategories(c *gin.Context) {

  categories, err := models.Category.FindAllCetegories(config.Server.DB)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "status": http.StatusNotFound,
    })
    return

  }
  c.JSON(http.StatusOK, gin.H{
    "status":   http.StatusOK,
    "response": categories,
  })
}
