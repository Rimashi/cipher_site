package handlers

import (
	"cipher_site/internal/ciphers"
	"cipher_site/internal/models"

	"github.com/gofiber/fiber/v2"
)

func Cipher(c *fiber.Ctx) error {
	id := c.Params("id")

	cipher, exists := ciphers.CiphersRegistry[id]
	if !exists {
		return c.Status(404).SendString("Шифр не найден")
	}

	// Преобразуем все зарегистрированные шифры в модели для меню
	var cipherModels []models.Cipher
	for id, cipher := range ciphers.CiphersRegistry {
		cipherModels = append(cipherModels, models.Cipher{
			ID:   id,
			Name: cipher.GetName(),
		})
	}

	hasModes := id == "Vigenere" // Определяем, поддерживает ли шифр режимы

	data := fiber.Map{
		"Title":    cipher.GetName(),
		"Cipher":   id,
		"Ciphers":  cipherModels,
		"HasKey":   cipher.RequiresKey(),
		"HasModes": hasModes, // Передаем информацию о поддержке режимов
	}

	return c.Render("cipher", data)
}
