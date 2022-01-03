package service

import "github.com/persiangophers/userland/internal/entity/models"

type IAccountService interface {
	SignUp(user models.User) (interface{}, error)
	Login(email string, password string) (string, error)
	EmailExist(email string) bool
}
