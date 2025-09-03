package models

import (
	"time"
)




type Role struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null;size:50" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	IsDefault   bool      `gorm:"default:false" json:"is_default"`
	CreatedAt   time.Time `json:"created_at"`
	
	// Relaciones
	Users       []User       `gorm:"many2many:user_roles;" json:"users,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

type UserRole struct {
	UserID     uint      `gorm:"primarykey" json:"user_id"`
	RoleID     uint      `gorm:"primarykey" json:"role_id"`
	AssignedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"assigned_at"`
	AssignedBy *uint     `json:"assigned_by,omitempty"` // Quién asignó el rol
}

// Tabla intermedia para roles-permisos
type RolePermission struct {
	RoleID       uint `gorm:"primarykey" json:"role_id"`
	PermissionID uint `gorm:"primarykey" json:"permission_id"`
}

func (u *User) HasRole(roleName string) bool {
	for _, role := range u.Roles {
		if role.Name == roleName {
			return true
		}
	}
	return false
}