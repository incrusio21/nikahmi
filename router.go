package main

import (
	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	// app.Group("/login", middleware.NonAuthMiddleware).Use("/", auth_page.Login)

	// auth := app.Group("", middleware.AuthMiddleware)
	// auth.Get("/logout", auth_page.Logout)
	// auth.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.SendString("Welcome")
	// })

}
