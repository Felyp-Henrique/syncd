package handlers

import "github.com/gofiber/fiber/v2"

func UsersHandlers(group fiber.Router) {
	group.Get("", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
}
