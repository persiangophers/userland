package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/pkg/logger/zap"
	"github.com/persiangophers/userland/internal/repository/postgres"
	"github.com/persiangophers/userland/internal/service/account"
	"github.com/persiangophers/userland/internal/transport/http/fiber"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
)

var serveCMD = &cli.Command{
	Name:    "serve",
	Aliases: []string{"s"},
	Usage:   "serve http",
	Action:  serve,
}

func serve(c *cli.Context) error {
	cfg := new(config.Config)
	config.ReadYAML("config.yaml", cfg)
	config.ReadEnv(cfg)

	f, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	logger := zap.New(f, zapcore.ErrorLevel)

	postgres, err := postgres.New(cfg.Postgres, logger)
	if err != nil {
		return err
	}

	accountService := account.New(cfg.Account, postgres, logger)
	server := fiber.New(logger, accountService)
	go func() {
		if err := server.Start(cfg.App.Address); err != nil {
			logger.Error(fmt.Sprintf("error happen while serving: %v", err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")

	if err := server.Shutdown(); err != nil {
		fmt.Println("\nRest server doesn't shutdown in 10 seconds")
	}

	return nil
}
