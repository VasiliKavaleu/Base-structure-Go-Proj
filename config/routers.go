package config

import (
  "goproj/controllers"
)

func (server *Server) InitializeRoutes() {
  r1 := server.Router.Group("/items")
  {
    r1.GET("/", controllers.GetAllItems)
  }

}
