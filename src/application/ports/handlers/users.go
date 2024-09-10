package handlers

import (
	"github.com/gofiber/fiber/v2"
)

const (
	UsersIndexHandlerRoute  string = "api.v1.users.index"
	UsersDetailHandlerRoute string = "api.v1.users.detail"
	UsersUpdateHandlerRoute string = "api.v1.users.update"
	UsersDeleteHandlerRoute string = "api.v1.users.delete"
)

func UsersIndexHandler(c *fiber.Ctx) error {
	return c.JSON(struct{}{})
}

func UsersDetailHandler(c *fiber.Ctx) error {
	return c.SendString(c.Params("id"))
}

func UsersUpdateHandler(c *fiber.Ctx) error {
	return c.SendString(c.Params("id"))
}

func UsersDeleteHandler(c *fiber.Ctx) error {
	return c.SendString(c.Params("id"))
}
