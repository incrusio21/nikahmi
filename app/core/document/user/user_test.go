package user

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/incrusio21/nikahmi/app"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// fiber instance
	app := app.Router
	app.Post("/user/create", Create)

	request := httptest.NewRequest(fiber.MethodPost, "/user/create", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

}
