package routes

import (
	"cipher_site/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// главная страница
	app.Get("/", handlers.Index)
	app.Get("/cipher/:id", handlers.Cipher)

	// API endpoints
	app.Post("/api/:id/encrypt", handlers.EncryptHandler)
	app.Post("/api/:id/decrypt", handlers.DecryptHandler)
}
