package login

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/app/core/model"
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

		auth := model.Auth{Name: request.Username, Fieldname: "password", Doctype: "User"}
		err = mysql.Db.First(&auth).Error
		if err == nil && utils.CheckPassword(auth.Password, request.Password) {
			sess := config.SetSession(ctx, []config.Session{{Name: "name", Value: request.Username}})

			sess.SetUser(request.Username)
			if err := sess.Save(auth.SimultaneousSessions); err == nil {
				ctx.Method(fiber.MethodGet)
				return ctx.Redirect("/")
			}

			login_err = fmt.Sprint(err)
		}

		if login_err == "" {
			login_err = "User atau Password yang digunakan salah"
		}
	}

	return ctx.Render("core/page/login/login", fiber.Map{
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
