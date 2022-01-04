package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type DbInstance struct {
	instance   *ClientInstance
	collection *mongo.Collection
}

func (db *DbInstance) DropDatabase() error {

	err := db.collection.Drop(context.TODO())
	if err != nil {
		db.instance.logger.Error("Unable to drop collection", zap.String("error", err.Error()))
		return err
	}

	return nil
}
