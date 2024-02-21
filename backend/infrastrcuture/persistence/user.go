package persistence

import (
	"backend/domain/model"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type userPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{db}
}

func (up *userPersistence) CreateUser(user *model.User) error {
	if err := up.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (up *userPersistence) UserByEmail(user *model.User, email string) error {
	if err := up.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}
