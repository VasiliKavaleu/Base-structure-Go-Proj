package routes

import (
  "github.com/gin-gonic/gin"
  "goproj/controllers"
)

var Router *gin.Engine

func Setup() {
  Router = gin.Default()
  items := Router.Group("/items")
  {
    items.GET("/", controllers.GetAllItems)
  }
  categories := Router.Group("/categories")
  {
    categories.GET("/", controllers.GetCategories)
  }

}
