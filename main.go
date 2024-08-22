package main

import (
	"log"
	"url-shortener/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	router.InitRoutes(app)
	app.Listen(":3000")
}