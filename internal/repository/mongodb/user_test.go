package mongodb

import (
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/entity/models"
	"github.com/persiangophers/userland/internal/repository"
	"github.com/persiangophers/userland/pkg/zap_logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"testing"
)

var (
	client   *mongo.Client
	database repository.IUserDb
	logger   *zap.Logger
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

func initTest() {
	cfg, _ := config.GetConfig("../../../config.yaml")

	logger, _ = zap_logger.InitLogger("../../../logs/test.log")
	defer zap_logger.SyncLogger(logger)

	client, _ := InitDatabase(cfg)
	database, _ = NewUserDatabase(client, logger)
}

func TestCreateUser(t *testing.T) {
	initTest()

	_, err := database.CreateUser(testUsers[0])
	if err != nil {
		t.Errorf("Unable to create the user: %s", err)
	}
}
