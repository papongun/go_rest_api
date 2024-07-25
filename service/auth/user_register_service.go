package auth

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	dto "github.com/papongun/go_todo/dto/auth"
	"github.com/papongun/go_todo/exception"
	"github.com/papongun/go_todo/repository"
)

// Interface
type AuthRegisterService interface {
	Register(request *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error)
}

// Singleton
var (
	authRegServiceOnce     sync.Once
	authRegServiceInstance AuthRegisterService
)

func GetAuthRegisterService(r repository.UserRepository) *AuthRegisterService {
	authRegServiceOnce.Do(func() {
		authRegServiceInstance = &AuthRegisterServiceImpl{r: r}
	})
	return &authRegServiceInstance
}

// Implement
type AuthRegisterServiceImpl struct {
	r repository.UserRepository
}

func (s *AuthRegisterServiceImpl) Register(request *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {

	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		return nil, exception.ValidationError{Message: err.Error()}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user, err := s.r.Save(request.Username, request.DisplayName, string(hashedPassword))

	if err != nil {
		return nil, exception.ValidationError{Message: err.Error()}
	}

	response := &dto.UserRegisterResponse{
		Username:    user.Username,
		DisplayName: user.DisplayName,
	}
	return response, nil
}
