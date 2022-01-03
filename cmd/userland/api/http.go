package api

import (
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/transport/http/fiber"
	"github.com/persiangophers/userland/pkg/zap_logger"
	"github.com/pkg/errors"
	"log"
)

func StartApi(port string) error {

	cfg, err := config.GetConfig("config.yaml")

	if err != nil {
		return errors.Errorf("Couldn't read config.yaml file: %s", err)
	}

	logger, err := zap_logger.InitLogger(cfg.Log.ApiLogPath)
	defer zap_logger.SyncLogger(logger)

	if err != nil {
		return errors.Errorf("Failed to initialize zap_logger: %s", err)
	}

	app := fiber.New(logger)

	if port == "" {
		if cfg.Api.Port == "" {
			log.Print("port wasn't defined in config.yaml file, using default port :3000")
			port = ":3000"
		} else {
			port = cfg.Api.Port
		}
	}

	return app.Start(port, cfg)
}
