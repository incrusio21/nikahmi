package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/db/mysql"
)

func Create(ctx *fiber.Ctx) {
	user := User{
		Name: "administrator",
	}

	response := mysql.Db.Create(&user)
	if response.Error != nil {
		panic(response.Error)
	}

	fmt.Println(response.RowsAffected)
}
