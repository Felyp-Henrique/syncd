package handlers

import (
	"Felyp-Henrique/syncd/src/application"

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

func init() {
	var (
		server          *fiber.App   = application.InjectHttpServer().GetServer()
		groupApiV1      fiber.Router = server.Group("/v1")
		groupApiV1Users fiber.Router = groupApiV1.Group("/users")
	)
	groupApiV1Users.Get("", UsersIndexHandler).Name(UsersIndexHandlerRoute)
	groupApiV1Users.Get("/:id", UsersDetailHandler).Name(UsersDetailHandlerRoute)
	groupApiV1Users.Put("/:id", UsersUpdateHandler).Name(UsersUpdateHandlerRoute)
	groupApiV1Users.Put("/:id", UsersDeleteHandler).Name(UsersDeleteHandlerRoute)
}
