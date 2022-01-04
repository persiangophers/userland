package user

import (
	"fmt"
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/entity/models"
	"github.com/persiangophers/userland/internal/repository/mongodb"
	"github.com/persiangophers/userland/internal/service/account"
	"github.com/persiangophers/userland/pkg/input"
	"github.com/persiangophers/userland/pkg/zap_logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
)

var (
	logger *zap.Logger
	cfg    *config.Config
)

func initConfigAndLogger() error {
	var err error

	cfg, err = config.GetConfig("config.yaml")

	if err != nil {
		return errors.Errorf("Couldn't read config.yaml file: %s", err)
	}

	logger, err = zap_logger.InitLogger(cfg.Log.ApiLogPath)
	defer zap_logger.SyncLogger(logger)

	if err != nil {
		return errors.Errorf("Failed to initialize zap_logger: %s", err)
	}

	return nil
}

func Add() error {

	err := initConfigAndLogger()

	if err != nil {
		return err
	}

	var firstName string = input.GetUserInput("First name", true, false)
	var lastName string = input.GetUserInput("Last name", true, false)
	var email string = input.GetUserInput("email", true, false)
	var password string = input.GetUserInput("secure_password", true, true)
	var birthDate string = input.GetUserInput("birth date", true, false)

	client, err := mongodb.InitClient(cfg, logger)

	defer client.DisconnectDatabase()

	if err != nil {
		logger.Error("Unable to connect to database", zap.String("error", err.Error()))
	}

	service, err := account.NewAccountService(client)

	if err != nil {
		logger.Error("Unable to load service", zap.String("error", err.Error()))
	}

	id, err := service.SignUp(models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		BirthDate: birthDate,
	})

	if err != nil {
		logger.Error("Unable sign up", zap.String("error", err.Error()))
	}

	logger.Info(fmt.Sprintf("User registered successfully with id %s", id))
	return nil
}

func Login() error {
	err := initConfigAndLogger()

	if err != nil {
		return err
	}

	var email string = input.GetUserInput("Email", true, false)
	var password string = input.GetUserInput("Password", true, true)

	client, err := mongodb.InitClient(cfg, logger)

	defer client.DisconnectDatabase()

	if err != nil {
		logger.Error("Unable to connect to database", zap.String("error", err.Error()))
	}

	service, err := account.NewAccountService(client)

	if err != nil {
		logger.Error("Unable to load service", zap.String("error", err.Error()))
	}

	token, err := service.Login(email, password)

	if err != nil {
		logger.Error(fmt.Sprintf("Unable login: %s", err))
	}

	log.Printf("you're logged in successfuly: %s", token)
	return nil
}
