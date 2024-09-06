package main

import (
	"Felyp-Henrique/syncd/src/application"
	"Felyp-Henrique/syncd/src/infrastructure/http"
)

func main() {
	app := http.New()
	application.HttpRoutes(app)
	app.Listen(":3000")
}
