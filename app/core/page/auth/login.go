package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/config"
)

func Login(ctx *fiber.Ctx) error {
	if ctx.Method() == fiber.MethodPost {
		sess := config.SetSession(ctx, "name", ctx.Query("name", "unknown user"))
		if err := sess.Save(); err != nil {
			panic(err)
		}

		return ctx.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	return ctx.Render("core/page/auth/login", fiber.Map{
		"Title": "Login",
		"Error": "",
	}, "views/login/master")
}

func Logout(ctx *fiber.Ctx) error {
	// Destry session
	if err := config.Session(ctx).Destroy(); err != nil {
		panic(err)
	}

	return ctx.Redirect("/", fiber.StatusTemporaryRedirect)
}
