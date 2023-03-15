package main

import (
	"github.com/ecom/pkg/config"
	"github.com/ecom/pkg/routes"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}
	config.SetUpDb()
}

func main() {
	routes.RunAPI()
}
