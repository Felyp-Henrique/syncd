package main

import (
	"Felyp-Henrique/syncd/src/application"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var (
		app *fiber.App = application.InjectHttpServer().GetServer()
	)
	app.Listen(":3000")
}
