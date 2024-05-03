package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/app/utils"
	"github.com/incrusio21/nikahmi/config"
	"github.com/incrusio21/nikahmi/db/mysql"
)

type LoginRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

func Login(ctx *fiber.Ctx) error {
	login_err := ""
	if ctx.Method() == fiber.MethodPost {
		request := new(LoginRequest)
		err := ctx.BodyParser(request)
		if err != nil {
			return err
		}

		user := Auth{Name: request.Username, Fieldname: "password", Doctype: "User"}

		err = mysql.Db.First(&user).Error
		if err == nil && utils.CheckPassword(user.Password, request.Password) {
			// ctx.Query("name", "unknown user")

			sess := config.SetSession(ctx, []config.Session{{Name: "name", Value: request.Username}})
			sess.SetUser(request.Username)
			if err := sess.Save(); err != nil {
				panic(err)
			}

			ctx.Method(fiber.MethodGet)
			return ctx.Redirect("/")
		}

		login_err = "User atau Password yang digunakan salah"
	}

	return ctx.Render("core/page/auth/login", fiber.Map{
		"Title": "Login",
		"Error": login_err,
	}, "views/login/master")
}

func Logout(ctx *fiber.Ctx) error {
	// Destry session
	if err := config.ReadSession(ctx).Destroy(); err != nil {
		panic(err)
	}

	return ctx.Redirect("/", fiber.StatusTemporaryRedirect)
}
