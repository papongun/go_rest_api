package repository

import (
	"gorm.io/gorm"

	"github.com/papongun/go_todo/entity"
)

type UserRepository interface {
	Save(username string, displayName string, password string) (*entity.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
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
