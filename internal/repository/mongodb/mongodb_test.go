package mongodb

import (
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/pkg/zap_logger"
	"testing"
)

func TestInitDatabase(t *testing.T) {

	cfg, err := config.GetConfig("../../../config.yaml")

	if err != nil {
		if err != nil {
			t.Errorf("Unable get config.yaml file: %s", err)
		}
	}

	logger, err := zap_logger.InitLogger("../../../logs/test.log")
	if err != nil {
		t.Errorf("Unable to initialize logger: %s", err)
	}

	defer zap_logger.SyncLogger(logger)

	client, err := InitDatabase(cfg)

	if err != nil {
		t.Errorf("Error initializing mongodb client: %s", err)
	}

	_, err = NewUserDatabase(client, logger)

	if err != nil {
		t.Errorf("Error initializing mongodb client: %s", err)
	}

	DisconnectDatabase(client, logger)

	if err != nil {
		t.Errorf("Error diconnecting mongodb client: %s", err)
	}
}
