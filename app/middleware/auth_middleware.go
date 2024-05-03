package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/config"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	sess := config.ReadSession(ctx)
	name := sess.Get("name")

	if name == nil {
		if err := sess.Destroy(); err != nil {
			panic(err)
		}

		return ctx.Redirect("/login", fiber.StatusTemporaryRedirect)
	}

	err := ctx.Next()

	return err
}

func NonAuthMiddleware(ctx *fiber.Ctx) error {
	sess := config.GetSession(ctx, "name")

	if sess != nil {
		return ctx.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	err := ctx.Next()

	return err
}
