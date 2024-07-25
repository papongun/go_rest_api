package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	dto "github.com/papongun/go_todo/dto/auth"
	"github.com/papongun/go_todo/entity"
	"github.com/papongun/go_todo/exception"
	s "github.com/papongun/go_todo/service/auth"
	repository_test "github.com/papongun/go_todo/test/repository"
)

func TestRegisterUserSuccess(t *testing.T) {
	mockRepo := new(repository_test.MockUserRepository)
	request := &dto.UserRegisterRequest{
		Username:    "testuser",
		DisplayName: "Test User",
		Password:    "password123",
	}
	expectedUser := &entity.User{
		Username:    request.Username,
		DisplayName: request.DisplayName,
		Password:    "hashedpassword",
	}
	mockRepo.On("Save", request.Username, request.DisplayName, mock.AnythingOfType("string")).Return(expectedUser, nil)
	service := s.AuthRegisterServiceImpl{R: mockRepo}

	response, err := service.Register(request)

	assert.NoError(t, err)
	assert.Equal(t, request.Username, response.Username)
	assert.Equal(t, request.DisplayName, response.DisplayName)
	mockRepo.AssertExpectations(t)
}

func TestRegisterUserMissingUsername(t *testing.T) {
	// TODO: Implement
}

func TestRegisterUserMissingDisplayname(t *testing.T) {
	// TODO: Implement
}

func TestRegisterUserMissingPassword(t *testing.T) {
	// TODO: Implement
}

func TestRegisterUserDuplicateUsername(t *testing.T) {
	mockRepo := new(repository_test.MockUserRepository)
	request := &dto.UserRegisterRequest{
		Username:    "duplicateuser",
		DisplayName: "Duplicate User",
		Password:    "password123",
	}
	mockRepo.On("Save", request.Username, request.DisplayName, mock.AnythingOfType("string")).Return((*entity.User)(nil), exception.ValidationError{Message: "duplicate key value violates unique constraint"})
	service := s.AuthRegisterServiceImpl{R: mockRepo}

	response, err := service.Register(request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "duplicate key value violates unique constraint", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestRegisterUserTooShortPassword(t *testing.T) {
	// TODO: Implement
}

func TestRegisterUserTooLongUsername(t *testing.T) {
	// TODO: Implement
}

func TestRegisterUserTooLongDisplayname(t *testing.T) {
	// TODO: Implement
}
