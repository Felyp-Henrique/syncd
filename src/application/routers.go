package application

import (
	"Felyp-Henrique/syncd/src/application/ports/handlers"

	"github.com/gofiber/fiber/v2"
)

func HttpRoutes(app *fiber.App) {
	var (
		groupApiV1      fiber.Router = nil
		groupApiV1Users fiber.Router = nil
	)

	// Group API V1
	//
	groupApiV1 = app.Group("/v1")

	// Group API V1 Users
	//
	groupApiV1Users = groupApiV1.Group("/users")
	handlers.UsersHandlers(groupApiV1Users)
}

func RpcRoutes() {

}
