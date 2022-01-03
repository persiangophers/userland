package fiber

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/transport/http/controllers"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @openApiFile  http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func (r *rest) routing(cfg *config.Config) {
	r.fiber.Post("account/login", func(ctx *fiber.Ctx) error {
		return controllers.Login(ctx, r.logger, cfg)
	})
	r.fiber.Post("account/signup", func(ctx *fiber.Ctx) error {
		return controllers.SignUp(ctx, r.logger, cfg)
	})
	r.fiber.Post("account/refresh", func(ctx *fiber.Ctx) error {
		return controllers.Refresh(ctx, r.logger, cfg)
	})

	r.fiber.Get("/swagger/*", swagger.Handler) // default
	r.fiber.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(map[string]string{
			"message": "Path not found",
		})
	})
}
