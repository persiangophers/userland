package mongodb

import (
	"github.com/persiangophers/userland/internal/entity/models"
	"testing"
)

var testUsers = []models.User{
	{
		FirstName: "Ali",
		LastName:  "Rasouli",
		Privileges: []string{
			"write",
		},
		Email:     "ali@gmail.com",
		Password:  "12345678",
		BirthDate: "1/2/3",
	},
}

func TestCreateUser(t *testing.T) {
	initTest()
	client, err := InitClient(cfg, logger)

	defer client.DisconnectDatabase()

	if err != nil {
		t.Errorf("Error initializing mongodb client: %s", err)
	}

	collection, err := client.NewDatabase("users")

	if err != nil {
		t.Errorf("Error initializing mongodb client: %s", err)
	}

	err = collection.DropDatabase()

	if err != nil {
		t.Errorf("Error dropping database: %s", err)
	}

	_, err = collection.CreateUser(testUsers[0])
	if err != nil {
		t.Errorf("Unable to create the user: %s", err)
	}
}
