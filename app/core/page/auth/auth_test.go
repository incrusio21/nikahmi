package auth

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	r_app "github.com/incrusio21/nikahmi/app"
	"github.com/incrusio21/nikahmi/app/middleware"
	"github.com/incrusio21/nikahmi/config"
	"github.com/incrusio21/nikahmi/db/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

// fiber instance
var cookie string
var app = r_app.Router

func Route() {
	app.Group("/login", middleware.NonAuthMiddleware).Use("/", Login)
	app.Get("/logout", Logout)
}

func TestSessionLogin(t *testing.T) {
	t.Run("LoginWithoutSession", TestLoginWithoutSession)
	t.Run("LoginPost", TestLoginPost)
	t.Run("LoginWithSession", TestLoginWithSession)

	t.Run("Logout", func(t *testing.T) {
		request := httptest.NewRequest(fiber.MethodGet, "/logout", nil)
		request.Header.Add("Cookie", "session_id="+cookie)

		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.Equal(t, 307, response.StatusCode)

		if config.Yaml.Session == "mysql" {
			var session string
			err := mysql.Db.Raw("select k from fiber_storage where k = ?", cookie).Scan(&session).Error
			assert.Nil(t, err)
			assert.Equal(t, "", session)
		}
	})
}

func TestLoginWithoutSession(t *testing.T) {
	Route()
	request := httptest.NewRequest(fiber.MethodGet, "/login", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)
}

func TestLoginPost(t *testing.T) {
	Route()
	body := strings.NewReader(`{"username": "Administrator", "password": "password"}`)
	request := httptest.NewRequest(fiber.MethodPost, "/login", body)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)
	require.NotNil(t, response.Header.Values("set-cookie"))
	assert.Equal(t, 307, response.StatusCode)

	cookie = getSessionID(response.Header.Values("set-cookie")[0])
}

func TestLoginWithSession(t *testing.T) {
	Route()
	request := httptest.NewRequest(fiber.MethodGet, "/login", nil)
	request.Header.Add("Cookie", "session_id="+cookie)
	response, err := app.Test(request)

	assert.Nil(t, err)
	assert.Equal(t, 307, response.StatusCode)
}
