package main

import (
		"goproj/config"
		"github.com/joho/godotenv"
		"log"
		"fmt"
		"os"
	)

var server = config.Server{}

func main() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("Getting values from env..")
	}


  configDB := config.ConfigDB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	fmt.Println(configDB)

	server.Initialize()
	server.Run(os.Getenv("IP_HOST"))

}
