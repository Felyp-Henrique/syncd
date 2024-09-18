package views

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var (
	//go:embed *
	embedStaticFolderFS embed.FS
)

func GetViewsHandler() fiber.Handler {
	return filesystem.New(filesystem.Config{
		Browse: true,
		Root:   http.FS(embedStaticFolderFS),
	})
}
