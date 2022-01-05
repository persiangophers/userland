package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/persiangophers/userland/internal/entiry/model"
	"github.com/persiangophers/userland/internal/transport/http/request"
	"github.com/persiangophers/userland/internal/transport/http/response"
)

func (f *FiberHandler) Login(c *fiber.Ctx) error {
	payload := &request.Login{}

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	token, err := f.account.Login(c.Context(), payload.Username, payload.Password)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(response.Error{
			Error: err.Error(),
		})
	}

	return c.JSON(&response.Login{
		Token: token,
	})
}

func (f *FiberHandler) Register(c *fiber.Ctx) error {
	payload := &request.Register{}

	if err := c.BodyParser(payload); err != nil {
		return err
	}

	err := f.account.Register(c.Context(), &model.User{
		Username:  payload.Username,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  payload.Password,
		Email:     payload.Email,
	})

	if err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return c.JSON(response.Error{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}
