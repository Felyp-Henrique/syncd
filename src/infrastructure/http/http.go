package server

import "github.com/gofiber/fiber/v2"

var (
	_app *fiber.App = nil
)

func GetHttpServer() *fiber.App {
	return _app
}

func init() {
	if _app == nil {
		_app = fiber.New()
	}
}
