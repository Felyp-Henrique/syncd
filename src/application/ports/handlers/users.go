package handlers

import "github.com/gofiber/fiber/v2"

func UsersIndex(c *fiber.Ctx) error {
	return c.JSON(struct{}{})
}
