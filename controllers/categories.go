package controllers

import (
  "github.com/gin-gonic/gin"
  "goproj/models"
  "goproj/config"
  "net/http"
)


func GetCategories(c *gin.Context) {
  category := models.Category{}

  categories, err := category.FindAllCategories(config.DB)

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "status": http.StatusNotFound,
    })
    return

  }
  c.JSON(http.StatusOK, gin.H{
    "categories": categories,
  })

}
