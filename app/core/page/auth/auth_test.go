package auth

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/app"
	"github.com/incrusio21/nikahmi/app/middleware"
	"github.com/incrusio21/nikahmi/config"
	"github.com/incrusio21/nikahmi/db/mysql"
	"github.com/stretchr/testify/assert"
)

func getSessionID(cookie string) string {
	parts := strings.Split(cookie, "; ")
	for _, part := range parts {
		if strings.HasPrefix(part, "session_id=") {
			return strings.TrimPrefix(part, "session_id=")
		}
	}
	return ""
}

func TestPostLogin(t *testing.T) {
	// fiber instance
	app := app.Router
	app.Group("/login", middleware.NonAuthMiddleware).Use("/", Login)
	app.Get("/logout", Logout)

	request := httptest.NewRequest(fiber.MethodGet, "/login", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	request = httptest.NewRequest(fiber.MethodPost, "/login", nil)
	response, err = app.Test(request)
	assert.Nil(t, err)
	assert.NotNil(t, response.Header.Values("set-cookie"))
	assert.Equal(t, 307, response.StatusCode)

	cookie := getSessionID(response.Header.Values("set-cookie")[0])

	request = httptest.NewRequest(fiber.MethodGet, "/login", nil)
	request.Header.Add("Cookie", "session_id="+cookie)
	response, err = app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, 307, response.StatusCode)

	request = httptest.NewRequest(fiber.MethodGet, "/logout", nil)
	request.Header.Add("Cookie", "session_id="+cookie)

	response, err = app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 307, response.StatusCode)

	if config.Yaml.Session == "mysql" {
		var session string
		err := mysql.Db.Raw("select k from fiber_storage where k = ?", cookie).Scan(&session).Error
		assert.Nil(t, err)
		assert.Equal(t, "", session)
	}
}
