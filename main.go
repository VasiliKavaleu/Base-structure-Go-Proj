package main

import (
		"goproj/config"
	)

var server = config.Server{}

func main() {

	server.Initialize()
	server.Run("0.0.0.0:8080")

}
