package repository

import "github.com/phucnh/go-app-sample/entity"

// UserRepository represent the user's repository
type UserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}
