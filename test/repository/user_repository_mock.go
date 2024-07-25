package repository_test

import (
	"github.com/stretchr/testify/mock"

	"github.com/papongun/go_todo/entity"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(username, displayName, password string) (*entity.User, error) {
	args := m.Called(username, displayName, password)
	return args.Get(0).(*entity.User), args.Error(1)
}
