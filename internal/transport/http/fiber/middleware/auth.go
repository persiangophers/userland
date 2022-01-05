package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/persiangophers/userland/pkg/myjwt"
)

func ValidateToken(c *fiber.Ctx) error {
	token := myjwt.FetchToken(c.Get("Authorization"))
	if token == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	ok, _ := myjwt.GetInstance().IsValid(token)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
