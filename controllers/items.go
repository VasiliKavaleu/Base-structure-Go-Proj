package controllers


import (
    "github.com/gin-gonic/gin"
)


func GetAllItems(c *gin.Context) {
  c.JSON(200, gin.H{
    "items": "many items",
  })
}
