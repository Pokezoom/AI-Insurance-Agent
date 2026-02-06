package service

import (
	"AI-Insurance-Agent/internal/middleware"
	"AI-Insurance-Agent/internal/model"
	"AI-Insurance-Agent/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(username, password, email string) (*model.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
		Role:     "agent",
		Status:   "active",
	}

	err := s.userRepo.Create(user)
	return user, err
}

func (s *AuthService) Login(username, password string) (string, *model.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", nil, errors.New("用户不存在")
	}

	if user.Status != "active" {
		return "", nil, errors.New("用户已被禁用")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("密码错误")
	}

	token, _ := s.generateToken(user)
	return token, user, nil
}

func (s *AuthService) generateToken(user *model.User) (string, error) {
	claims := middleware.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(middleware.JWTSecret)
}
