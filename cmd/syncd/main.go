package main

import (
	"Felyp-Henrique/syncd/src/presentations/views"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use("/", views.GetViewsHandler())
	app.Listen(":8080")
}
