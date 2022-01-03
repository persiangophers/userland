package mongodb

import (
	"context"
	"fmt"
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
)

type userDb struct {
	db     *mongo.Collection
	logger *zap.Logger
}

func InitDatabase(cfg *config.Config) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=%s&authMechanism=%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.AuthDatabase,
		cfg.Database.AuthMechanism))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err.Error())
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewUserDatabase(client *mongo.Client, logger *zap.Logger) (repository.IUserDb, error) {
	collection := client.Database("userland").Collection("users")
	return &userDb{db: collection, logger: logger}, nil
}

func DisconnectDatabase(client *mongo.Client, logger *zap.Logger) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		logger.Error("Unable to disconnect from database", zap.String("error", err.Error()))
	}
}

func DropUserDatabase(client *mongo.Client) error {
	collection := client.Database("userland").Collection("users")

	err := collection.Drop(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
