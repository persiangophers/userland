package repository

import "github.com/persiangophers/userland/internal/entity/models"

type IUserDb interface {
	CreateUser(user models.User) (interface{}, error)
	UpdateUser(user models.User) (models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	DeleteUser(id string) error
}
