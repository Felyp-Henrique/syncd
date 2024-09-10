//go:build wireinject
// +build wireinject

package application

import (
	"Felyp-Henrique/syncd/src/infrastructure/server"

	"github.com/google/wire"
)

func InjectHttpServer() *server.HttpServer {
	wire.Build(server.NewHttpServer)
	return &server.HttpServer{}
}
