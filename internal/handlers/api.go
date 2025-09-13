package handlers

import (
	"cipher_site/internal/ciphers"
	"log"

	"github.com/gofiber/fiber/v2"
)

type CipherRequest struct {
	Text string `json:"text"`
	Key  string `json:"key,omitempty"`
	Mode string `json:"mode,omitempty"`
}

func EncryptHandler(c *fiber.Ctx) error {
	return processCipherRequest(c, true)
}

func DecryptHandler(c *fiber.Ctx) error {
	return processCipherRequest(c, false)
}

func processCipherRequest(c *fiber.Ctx, encrypt bool) error {
	cipherID := c.Params("id")
	var req CipherRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат запроса"})
	}

	cipher, exists := ciphers.CiphersRegistry[cipherID]
	if !exists {
		return c.Status(404).JSON(fiber.Map{"error": "Шифр не найден"})
	}

	if vigenereCipher, ok := cipher.(interface{ SetMode(mode string) }); ok && req.Mode != "" {
		vigenereCipher.SetMode(req.Mode)
	}

	var result string
	var err error

	if encrypt {
		result, err = cipher.Encrypt(req.Text, req.Key)
	} else {
		result, err = cipher.Decrypt(req.Text, req.Key)
	}

	if err != nil {
		log.Printf("Ошибка обработки: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"result": result})
}
