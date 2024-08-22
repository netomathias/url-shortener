package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"url-shortener/models"
	"url-shortener/repositories"

	"github.com/gofiber/fiber/v2"
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

	hash := sha256.New()
	hash.Write([]byte(urlShortenerModel.OriginalUrl))
	alias := base64.URLEncoding.EncodeToString(hash.Sum(nil))[:4]

	urlShortenerModel.Alias = alias
	urlShortenerModel.ShortenedUrl = "http://localhost:3000/" + alias
	
	err := r.UrlRepo.Save(urlShortenerModel)
	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
		
	}

	return c.JSON(fiber.Map{"message": "success"})
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
	
	return c.Redirect(result.OriginalUrl, 301)
}
