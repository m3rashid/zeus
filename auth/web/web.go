package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var (
	//go:embed dist/*
	index embed.FS
)

var uiRoutes = []string{"/hello"}

func HandleWeb(app *fiber.App) {
	_index, err := fs.Sub(index, "dist")
	if err != nil {
		panic(err)
	}

	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(_index),
		Index:  "index.html",
		Browse: false,
	}))

	for _, path := range uiRoutes {
		app.Get(path, func(ctx *fiber.Ctx) error {
			return filesystem.SendFile(ctx, http.FS(_index), "index.html")
		})
	}
}
