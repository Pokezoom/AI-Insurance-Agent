package service

import (
	"AI-Insurance-Agent/internal/model"
	"AI-Insurance-Agent/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetProfile(userID int64) (*model.User, error) {
	return s.userRepo.FindByID(userID)
}

func (s *UserService) ChangePassword(userID int64, oldPassword, newPassword string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("原密码错误")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	return s.userRepo.UpdatePassword(userID, string(hashedPassword))
}

func (s *UserService) ListUsers(page, pageSize int) ([]model.User, int64, error) {
	return s.userRepo.List(page, pageSize)
}

func (s *UserService) UpdateUserStatus(userID int64, status string) error {
	return s.userRepo.UpdateStatus(userID, status)
}
