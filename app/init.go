package app

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/incrusio21/nikahmi/config"
)

//go:embed *
var Templatesfs embed.FS

var Router *fiber.App

func init() {
	conf, err := config.Read()
	if err != nil {
		panic(err)
	}

	engine := html.NewFileSystem(http.FS(Templatesfs), ".html")
	engine.AddFunc(
		// add unescape function
		"WEBSITE_NAME", func() string { return conf.App.Name },
	)

	Router = fiber.New(fiber.Config{
		// Prefork:      true,
		Views:        engine,
		ErrorHandler: ErrorHandler,
	})

	Router.Static("/public", "./source")
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// ctx.Status(fiber.StatusInternalServerError)
	return ctx.Render("views/errors/index", "")
}
