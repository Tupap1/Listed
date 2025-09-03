package models

import (
	"time"
)



type Permission struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null;size:100" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Resource    string    `gorm:"not null;size:50" json:"resource"` // products, sales, inventory, etc.
	Action      string    `gorm:"not null;size:50" json:"action"`   // read, create, update, delete, admin
	CreatedAt   time.Time `json:"created_at"`
	
	Roles []Role `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}


func (u *User) HasPermission(permission string) bool {
	for _, role := range u.Roles {
		for _, perm := range role.Permissions {
			if perm.Name == permission {
				return true
			}
		}
	}
	return false
}