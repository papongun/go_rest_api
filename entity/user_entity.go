package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex; not null"`
	DisplayName string `gorm:"not null; "`
	Password    string `gorm:"not null"`
}
