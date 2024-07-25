package auth_test

import (
	"github.com/stretchr/testify/mock"

	"github.com/papongun/go_todo/dto/auth"
)

type MockAuthRegisterService struct {
	mock.Mock
}

func (m *MockAuthRegisterService) Register(request *auth.UserRegisterRequest) (*auth.UserRegisterResponse, error) {
	args := m.Called(request)
	return args.Get(0).(*auth.UserRegisterResponse), args.Error(1)
}
