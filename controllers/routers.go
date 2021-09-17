package controllers

import (
  "goproj/config"
)

func InitializeRoutes((server *Server) ) {
  items := server.Router.Group("/items")
  {
    items.GET("/", controllers.GetAllItems)
  }
  categories := server.Router.Group("/categories")
  {
    categories.GET("/", controllers.GetCategories)
  }

}
