package router

import (
	"os"
	"url-shortener/database"
	"url-shortener/handlers"
	"url-shortener/repositories"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	database, _ := database.NewMongoDB(&database.MongoConfig{URI: os.Getenv("MONGO_URI"), Database: "url-shortener"})
	repository := repositories.NewUrlRepositoryImpl(database.Database)
	handlers := handlers.NewUrlShortenerHandler(repository)
	app.Get("/:alias", handlers.ResolveURL)
	app.Post("/", handlers.ShortenURL)
}