package routes

import (
	"goproj/controllers"

	"github.com/gin-gonic/gin"
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
		categories.POST("/", controllers.CreateCategory)
		categories.GET("/:id", controllers.GetCategory)
		categories.DELETE("/:id", controllers.DeleteCategory)
		categories.PUT("/:id", controllers.UpdateCategory)
	}

}
