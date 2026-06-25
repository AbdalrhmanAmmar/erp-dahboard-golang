package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
Email        string `json:"email" gorm:"uniqueIndex;not null"`
	Username     string `json:"username" gorm:"uniqueIndex;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	Roles        []Role `json:"roles" gorm:"many2many:user_roles;"`

}

type Role struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `json:"-" gorm:"many2many:user_roles;"`

}