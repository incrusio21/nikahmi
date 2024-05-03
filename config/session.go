package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/app/middleware/session"
)

func ReadSession(c *fiber.Ctx) *session.Session {
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	return sess
}

func GetSession(c *fiber.Ctx, session_name string) any {
	return ReadSession(c).Get(session_name)
}

func SetSession(c *fiber.Ctx, session []Session) *session.Session {
	sess := ReadSession(c)

	for _, val := range session {
		sess.Set(val.Name, val.Value)
	}

	return sess
}

func DeleteSession(c *fiber.Ctx, session_name string) *session.Session {
	sess := ReadSession(c)
	sess.Delete(session_name)

	return sess

}

// func SaveSession(c *fiber.Ctx) {
// 	ReadSession(c).Save()
// }
