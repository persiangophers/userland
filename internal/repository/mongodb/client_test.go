package mongodb

import (
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/pkg/zap_logger"
	"go.uber.org/zap"
	"log"
	"testing"
)

var (
	logger *zap.Logger
	cfg    *config.Config
)

func initTest() {
	var err error
	cfg, err = config.GetConfig("../../../config.yaml")

	if err != nil {
		if err != nil {
			log.Fatalf("Unable get config.yaml file: %s", err)
		}
	}

	logger, err = zap_logger.InitLogger("../../../logs/test.log")
	defer zap_logger.SyncLogger(logger)
	if err != nil {
		log.Fatalf("Unable to initialize logger: %s", err)
	}
}

func TestInitClient(t *testing.T) {
	initTest()
	client, err := InitClient(cfg, logger)

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

	err = client.DisconnectDatabase()
	if err != nil {
		t.Errorf("Error diconnecting mongodb client: %s", err)
	}

}
