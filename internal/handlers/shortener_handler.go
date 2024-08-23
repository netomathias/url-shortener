package handlers

import (
	"os"
	"url-shortener/internal/models"
	"url-shortener/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/rand"
)


type UrlShortenerHandler struct {
	UrlRepo repositories.UrlRepository
}

func NewUrlShortenerHandler(UrlRepo repositories.UrlRepository) UrlShortenerHandler {
	return UrlShortenerHandler{
		UrlRepo: UrlRepo,
	}
}

func (r *UrlShortenerHandler) ShortenURL(c *fiber.Ctx) error {
	c.Request().Header.Add("Content-Type", "application/json")
	urlShortenerModel := models.NewUrlShortenerModel()
	if err := c.BodyParser(&urlShortenerModel); err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	for {
		alias := generateAlias()
		_, err := r.UrlRepo.FindByAlias(alias)
		if err != nil {
			urlShortenerModel.Alias = alias
			urlShortenerModel.ShortUrl = os.Getenv("HOSTNAME") + alias
			break
		}
	}

	err := r.UrlRepo.Save(urlShortenerModel)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
		
	}

	return c.JSON(fiber.Map{"short_url": urlShortenerModel.ShortUrl})
}

func (r *UrlShortenerHandler) ResolveURL(c *fiber.Ctx) error {
	c.Response().Header.Add("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
	c.Response().Header.Add("Pragma", "no-cache")
	c.Response().Header.Add("Expires", "0")
	
	alias := c.Params("alias")
	result, err := r.UrlRepo.FindByAlias(alias)
	if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short-url not found in db"})
    }

	err = r.UrlRepo.UpdateClicks(alias)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	
	return c.Redirect(result.LongUrl, 301)
}

func generateAlias() string {
	const aliasLength = 4
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, aliasLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
