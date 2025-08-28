package services

import (
	"errors"
	"os"
	"time"

	"auth-service/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	
	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *AuthService) LoginUser(email, password string) (*models.User, error) {
	var user models.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("invalid credentials")
	}
	
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}

func (s *AuthService) GenerateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET is not set")
	}
	
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthService) GenerateRefreshToken(userID uint) (*models.RefreshToken, error) {
	token := &models.RefreshToken{
		UserID:    userID,
		Token:     "random_string_123", 
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}
	result := s.db.Create(token)
	if result.Error != nil {
		return nil, result.Error
	}
	return token, nil
}
