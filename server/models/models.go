package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}

func (User) TableName() string {
	return "users"
}
