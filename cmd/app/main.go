package main

import (
	"log"

	"github.com/angelmendozacap/go-structure/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := initRoutes()

	log.Fatal(e.Start(config.GetPort()))
}
