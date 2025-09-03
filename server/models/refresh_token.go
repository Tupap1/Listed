package models

import (
	"time"
)

type RefreshToken struct {
	ID          uint       `gorm:"primarykey" json:"id"`
	UserID      uint       `gorm:"not null" json:"user_id"`
	TokenHash   string     `gorm:"not null;size:255;index" json:"-"`
	DeviceID    string     `gorm:"size:255" json:"device_id"`
	ExpiresAt   time.Time  `gorm:"not null" json:"expires_at"`
	CreatedAt   time.Time  `json:"created_at"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	
	ParentTokenID *uint `json:"parent_token_id,omitempty"`
	
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (rt *RefreshToken) IsRevoked() bool {
	return rt.RevokedAt != nil
}

func (rt *RefreshToken) IsExpired() bool {
	return time.Now().After(rt.ExpiresAt)
}

func (rt *RefreshToken) IsValid() bool {
	return !rt.IsRevoked() && !rt.IsExpired()
}