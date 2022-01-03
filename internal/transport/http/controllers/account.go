package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/entity/models"
	"github.com/persiangophers/userland/internal/repository/mongodb"
	"github.com/persiangophers/userland/internal/service/account"
	"github.com/persiangophers/userland/internal/transport/http/request_models"
	"github.com/persiangophers/userland/pkg/fiber_validator"
	"github.com/persiangophers/userland/pkg/jwt_generator"
	"go.uber.org/zap"
)

// Login godoc
//// @Summary Login to your account
//// @Accept  json
//// @Produce  json
//// @Param account  body request_models.Login  true  "Login to Account"
//// @Success 200 {object} response_models.Token
//// @Router /account/login [post]
func Login(c *fiber.Ctx, logger *zap.Logger, cfg *config.Config) error {

	model := new(request_models.Login)

	if isValid, err := fiber_validator.Validate(model, c); !isValid {
		return err
	}

	client, err := mongodb.InitDatabase(cfg)

	defer mongodb.DisconnectDatabase(client, logger)

	if err != nil {
		logger.Error("Unable to connect to database: %s", zap.String("error", err.Error()))
		return c.Status(500).JSON(fiber.Map{
			"message": "Something unexpected happened. Contact developer to fix it",
		})
	}

	service, err := account.NewAccountService(client, logger)

	if err != nil {
		logger.Error("Unable to get user service", zap.String("error", err.Error()))
		return c.Status(500).JSON(fiber.Map{
			"message": "Something unexpected happened. Contact developer to fix it",
		})
	}

	userId, err := service.Login(model.Email, model.Password)
	if err != nil {
		logger.Error("Unable to sign up user", zap.String("error", err.Error()))
		return c.Status(400).JSON(fiber.Map{
			"message": "Something unexpected happened, notify developer to fix it",
		})
	}

	token, err := jwt_generator.GenerateToken(userId, cfg)
	if err != nil {
		logger.Error("Unable to generate token", zap.String("error", err.Error()))
		return c.Status(500).JSON(fiber.Map{
			"message": "Something unexpected happened, notify developer to fix it",
		})
	}

	return c.JSON(token)
}

// SignUp godoc
//// @Summary SignUp to your account
//// @Accept  json
//// @Produce  json
//// @Param account formData request_models.SignUp  true  "Create Account"
//// @Router /account/signup [post]
func SignUp(c *fiber.Ctx, logger *zap.Logger, cfg *config.Config) error {

	model := new(request_models.SignUp)

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(406).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = c.SaveFile(file, fmt.Sprintf("./avatars/%s", file.Filename))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if isValid, err := fiber_validator.Validate(model, c); !isValid {
		return err
	}

	client, err := mongodb.InitDatabase(cfg)

	defer mongodb.DisconnectDatabase(client, logger)

	if err != nil {
		logger.Error("Unable to connect to database", zap.String("error", err.Error()))
		return c.Status(500).JSON(fiber.Map{
			"message": "Something unexpected happened. Contact developer to fix it",
		})
	}

	service, err := account.NewAccountService(client, logger)

	if err != nil {
		logger.Error("Unable to get user service", zap.String("error", err.Error()))
		return c.Status(500).JSON(fiber.Map{
			"message": "Something unexpected happened. Contact developer to fix it",
		})
	}

	if service.EmailExist(model.Email) {
		return c.Status(406).JSON(fiber.Map{
			"message": "email already exists",
		})
	}

	id, err := service.SignUp(models.User{
		FirstName:  model.FirstName,
		LastName:   model.LastName,
		Email:      model.Email,
		Password:   model.Password,
		BirthDate:  model.BirthDate,
		Privileges: model.Privileges,
	})

	if err != nil {
		logger.Error("Unable to sign up user", zap.String("error", err.Error()))
		return c.Status(500).JSON(fiber.Map{
			"message": "Something unexpected happened. Contact developer to fix it",
		})
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}

// Refresh godoc
//// @Summary Refresh your tokens
//// @Accept  json
//// @Produce  json
//// @Param account formData request_models.Refresh  true  "Refresh Token"
//// @Router /account/refresh [post]
func Refresh(c *fiber.Ctx, logger *zap.Logger, cfg *config.Config) error {
	model := new(request_models.Refresh)

	if isValid, err := fiber_validator.Validate(model, c); !isValid {
		return err
	}

	token, err := jwt.Parse(model.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Jwt.Secret), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid or expired token",
		})
	}

	newToken, err := jwt_generator.GenerateToken(claims["id"].(string), cfg)
	if err != nil {
		logger.Error("Unable to generate jwt token", zap.String("error", err.Error()))
		return c.Status(500).JSON(fiber.Map{
			"message": "Something unexpected happened. Notify developer to fix it",
		})
	}

	return c.JSON(newToken)
}
