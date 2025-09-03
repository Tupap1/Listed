package models

import (
	"fmt"
	"time"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Username    string    `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Email       string    `gorm:"uniqueIndex;not null;size:255" json:"email"`
	PasswordHash string   `gorm:"not null;size:255" json:"-"` 
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	ProfilePicture *string `gorm:"size:255" json:"profile_picture,omitempty"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	
	Roles        []Role        `gorm:"many2many:user_roles;" json:"roles,omitempty"`
	RefreshTokens []RefreshToken `gorm:"foreignKey:UserID" json:"-"`
	
}


func CreateUser(db *gorm.DB, username, email, password string, roles []Role) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error al hashear la contraseña: %w", err)
	}

	user := &User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		IsActive:     true,
	}

	if roles != nil {
		user.Roles = roles
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("error al crear el usuario en la base de datos: %w", err)
	}

	return user, nil
}


func (u *User) BeforeCreate(tx *gorm.DB) error {
	if len(u.Roles) == 0 {
		var defaultRole Role
		if err := tx.Where("is_default = ?", true).First(&defaultRole).Error; err == nil {
			u.Roles = []Role{defaultRole}
		}
	}
	return nil
}


func (u *User) GetAllPermissions() []string {
	permissionSet := make(map[string]bool)
	permissions := []string{}
	
	for _, role := range u.Roles {
		for _, perm := range role.Permissions {
			if !permissionSet[perm.Name] {
				permissionSet[perm.Name] = true
				permissions = append(permissions, perm.Name)
			}
		}
	}
	
	return permissions
}

