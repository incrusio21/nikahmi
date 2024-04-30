package app

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var Templatesfs embed.FS

var engine = html.NewFileSystem(http.FS(Templatesfs), ".html")

var App = fiber.New(fiber.Config{
	Prefork:      true,
	Views:        engine,
	ErrorHandler: ErrorHandler,
})

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// ctx.Status(fiber.StatusInternalServerError)
	return ctx.Render("views/errors/index", "")
}
