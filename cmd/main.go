package main

import (
	// Этот импорт активирует регистрацию шифров
	"cipher_site/internal/routes"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// движок шаблонов (html/template)
	// путь к шаблонам относительно текущего файла main.go
	basePath, err := filepath.Abs(filepath.Join(".", "internal", "templates"))
	if err != nil {
		log.Fatal(err)
	}

	engine := html.New(basePath, ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// подключаем маршруты
	routes.Setup(app)

	// статика (css/js)
	app.Static("/static", "./static")

	log.Println("Server is running on http://localhost:8000")
	log.Fatal(app.Listen(":8000"))
}
