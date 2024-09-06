package main

import "Felyp-Henrique/syncd/src/infrastructure/http"

func main() {
	app := http.New()
	app.Server()
}
