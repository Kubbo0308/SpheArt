package repository

import "backend/domain/model"

type UserRepository interface {
	CreateUser(user *model.User) error
	UserByEmail(user *model.User, email string) error
}
