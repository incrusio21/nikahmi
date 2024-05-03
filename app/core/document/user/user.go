package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/app/core/page/auth"
	"github.com/incrusio21/nikahmi/app/utils"
	"github.com/incrusio21/nikahmi/db/mysql"
)

func Create(ctx *fiber.Ctx) error {
	pass, err := utils.HashPassword("password")
	if err != nil {
		panic(err)
	}

	user := &User{
		Name: "Administrator",
		Auth: auth.Auth{Password: pass, Fieldname: "password"},
	}

	response := mysql.Db.Create(user)
	if response.Error != nil {
		panic(response.Error)
	}

	fmt.Println(response.RowsAffected)

	return ctx.SendString("Berhasil Di Input")
}
