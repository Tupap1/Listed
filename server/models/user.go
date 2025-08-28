package models


import (
	"gorm.io/gorm"
)

type User struct {
	ID 			 uint	`gorm:"not null;primaryKey"`
	RoleID		 uint 	`gorm:"not null;"`
	Username 	 string
	Email        string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`             
	Name         string
}

type RefreshToken struct {
	gorm.Model
	Token string `gorm:"not null"`
	UserID uint   
}
