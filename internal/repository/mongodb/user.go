package mongodb

import (
	"context"
	"github.com/persiangophers/userland/internal/entity/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (userDb *userDb) CreateUser(user models.User) (interface{}, error) {
	result, err := userDb.db.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (userDb *userDb) UpdateUser(user models.User) (models.User, error) {
	return user, nil
}

func (userDb *userDb) GetUserById(id string) (models.User, error) {
	return models.User{}, nil
}

func (userDb *userDb) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := userDb.db.FindOne(context.TODO(), bson.M{
		"email": email,
	}).Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (userDb *userDb) DeleteUser(id string) error {
	return nil
}
