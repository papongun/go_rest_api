package auth_test

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	c "github.com/papongun/go_todo/controller/auth"
	"github.com/papongun/go_todo/dto/auth"
	dto "github.com/papongun/go_todo/dto/auth"
	"github.com/papongun/go_todo/exception"
	mock "github.com/papongun/go_todo/test/service/auth"
)

func TestRegisterSuccess(t *testing.T) {
	app := fiber.New()
	mockService := new(mock.MockAuthRegisterService)
	controller := c.UserRegisterContoller{S: mockService}
	request := &dto.UserRegisterRequest{
		Username:    "testuser",
		DisplayName: "Test User",
		Password:    "password123",
	}
	expectedResponse := &dto.UserRegisterResponse{
		Username:    request.Username,
		DisplayName: request.DisplayName,
	}
	mockService.On("Register", request).Return(expectedResponse, nil)

	app.Post("/register", controller.Register)
	req := httptest.NewRequest("POST", "/register", strings.NewReader(`{
		"username": "testuser",
		"displayName": "Test User",
		"password": "password123"
	}`))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "Register success", body["message"])
	data, ok := body["data"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "testuser", data["username"])
	assert.Equal(t, "Test User", data["displayName"])

	mockService.AssertExpectations(t)
}

func TestRegisterFailure(t *testing.T) {
	app := fiber.New()
	mockService := new(mock.MockAuthRegisterService)
	controller := c.UserRegisterContoller{S: mockService}

	request := &auth.UserRegisterRequest{
		Username:    "testuser",
		DisplayName: "Test User",
		Password:    "password123",
	}

	// Set up the mock expectations
	mockService.On("Register", request).Return((*dto.UserRegisterResponse)(nil), exception.ValidationError{Message: "Registration failed"})

	// Set up a test route and send a request
	app.Post("/register", controller.Register)
	req := httptest.NewRequest("POST", "/register", strings.NewReader(`{
		"username": "testuser",
		"displayName": "Test User",
		"password": "password123"
	}`))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, 400, resp.StatusCode)

	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, "Field validation error", body["error"])

	// Assert that expectations were met
	mockService.AssertExpectations(t)
}
