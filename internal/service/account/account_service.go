package account

import (
	"errors"
	"github.com/persiangophers/userland/internal/entity/models"
	"github.com/persiangophers/userland/internal/repository"
	"github.com/persiangophers/userland/internal/repository/mongodb"
	"github.com/persiangophers/userland/internal/service"
	"github.com/persiangophers/userland/pkg/secure_password"
)

type accountService struct {
	db repository.IUserDb
}

func NewAccountService(client *mongodb.ClientInstance) (service.IAccountService, error) {
	collection, err := client.NewDatabase("users")

	if err != nil {
		return nil, err
	}

	return &accountService{
		db: collection,
	}, nil

}

func (service *accountService) SignUp(user models.User) (interface{}, error) {

	password := secure_password.HashAndSalt([]byte(user.Password))
	user.Password = password
	id, err := service.db.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (service *accountService) Login(email string, password string) (string, error) {

	user, err := service.db.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	match := secure_password.ComparePasswords(user.Password, password)

	if !match {
		return "", errors.New("secure_password doesn't match")
	}

	return user.Id.Hex(), nil
}

func (service *accountService) EmailExist(email string) bool {
	_, err := service.db.GetUserByEmail(email)

	if err == nil {
		return true
	}
	return false
}
