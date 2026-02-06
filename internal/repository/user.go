package repository

import (
	"AI-Insurance-Agent/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) UpdatePassword(id int64, newPassword string) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("password", newPassword).Error
}

func (r *UserRepository) List(page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	offset := (page - 1) * pageSize
	err := r.db.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(pageSize).Find(&users).Error
	return users, total, err
}

func (r *UserRepository) UpdateStatus(id int64, status string) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("status", status).Error
}
