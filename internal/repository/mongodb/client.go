package mongodb

import (
	"context"
	"fmt"
	"github.com/persiangophers/userland/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type ClientInstance struct {
	client   *mongo.Client
	database *mongo.Database
	logger   *zap.Logger
}

func InitClient(cfg *config.Config, logger *zap.Logger) (*ClientInstance, error) {

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
		logger.Error("Unable to connect to database: %s", zap.String("error", err.Error()))
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logger.Error("Unable to ping database: %s", zap.String("error", err.Error()))
		return nil, err
	}

	return &ClientInstance{
		client:   client,
		database: client.Database(cfg.Database.Name),
		logger:   logger,
	}, nil
}

func (client *ClientInstance) NewDatabase(collection string) (*DbInstance, error) {
	return &DbInstance{collection: client.database.Collection(collection), instance: client}, nil
}

func (client *ClientInstance) DisconnectDatabase() error {
	err := client.client.Disconnect(context.TODO())
	if err != nil {
		client.logger.Error("Unable to disconnect from database", zap.String("error", err.Error()))
		return err
	}

	return nil
}
