package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/app/core/page/login"
	"github.com/incrusio21/nikahmi/app/middleware"
)

func Routing(app *fiber.App) {
	app.Group("/login", middleware.NonAuthMiddleware).Use("/", login.Login)

	auth := app.Group("", middleware.AuthMiddleware)
	auth.Get("/logout", login.Logout)
	auth.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome")
	})

}
