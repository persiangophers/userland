package mongodb

import (
	"context"
	"github.com/persiangophers/userland/internal/entity/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DbInstance) CreateUser(user models.User) (interface{}, error) {
	result, err := db.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (db *DbInstance) UpdateUser(user models.User) (models.User, error) {
	return user, nil
}

func (db *DbInstance) GetUserById(id string) (models.User, error) {
	return models.User{}, nil
}

func (db *DbInstance) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := db.collection.FindOne(context.TODO(), bson.M{
		"email": email,
	}).Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (db *DbInstance) DeleteUser(id string) error {
	return nil
}
