package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Session(c *fiber.Ctx) *session.Session {
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	return sess
}

func GetSession(c *fiber.Ctx, session_name string) any {
	return Session(c).Get(session_name)
}

func SetSession(c *fiber.Ctx, session_name string, value any) *session.Session {
	sess := Session(c)
	sess.Set(session_name, value)

	return sess
}

func DeleteSession(c *fiber.Ctx, session_name string) *session.Session {
	sess := Session(c)
	sess.Delete(session_name)

	return sess

}

func SaveSession(c *fiber.Ctx) {
	Session(c).Save()
}
