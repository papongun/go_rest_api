package repository

import (
	"sync"

	"gorm.io/gorm"

	"github.com/papongun/go_todo/config"
	"github.com/papongun/go_todo/entity"
)

// Interface
type UserRepository interface {
	Save(username string, displayName string, password string) (*entity.User, error)
}

// Instance
var (
	userRepoOnce     sync.Once
	userRepoInstance UserRepository
)

func GetUserRepository() UserRepository {
	userRepoOnce.Do(func() {
		userRepoInstance = &UserRepositoryImpl{Db: config.GetDatabase()}
	})
	return userRepoInstance
}

// Implement
type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (r *UserRepositoryImpl) Save(username string, displayName string, password string) (*entity.User, error) {
	user := entity.User{
		Username:    username,
		DisplayName: displayName,
		Password:    password,
	}

	tx := r.Db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
