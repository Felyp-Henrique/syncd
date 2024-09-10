package server

import "github.com/gofiber/fiber/v2"

var (
	_fiberServer *fiber.App = fiber.New()
)

type HttpServer struct {
	_server *fiber.App
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		_server: _fiberServer,
	}
}

func (h *HttpServer) GetServer() *fiber.App {
	return h._server
}
