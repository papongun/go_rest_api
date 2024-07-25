package repository

import (
	"sync"

	"gorm.io/gorm"

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

func GetUserRepository(db *gorm.DB) *UserRepository {
	userRepoOnce.Do(func() {
		userRepoInstance = &UserRepositoryImpl{db: db}
	})
	return &userRepoInstance
}

// Implement
type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) Save(username string, displayName string, password string) (*entity.User, error) {
	user := entity.User{
		Username:    username,
		DisplayName: displayName,
		Password:    password,
	}

	tx := r.db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
