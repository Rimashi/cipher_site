package handlers

import (
	"cipher_site/internal/ciphers"
	"cipher_site/internal/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	// Преобразуем все зарегистрированные шифры в модели для меню
	var cipherModels []models.Cipher
	for id, cipher := range ciphers.CiphersRegistry {
		cipherModels = append(cipherModels, models.Cipher{
			ID:   id,
			Name: cipher.GetName(),
		})
	}

	data := fiber.Map{
		"Title":   "Cyphers - Главная",
		"Ciphers": cipherModels,
	}
	return c.Render("index", data)
}
