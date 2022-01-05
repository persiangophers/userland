package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/persiangophers/userland/internal/transport/http/fiber/middleware"
)

func (h *httpHandler) setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	h.setupPublicRoutes(api)
	h.setupProtectedRoutes(api)
}

func (h *httpHandler) setupPublicRoutes(api fiber.Router) {
	api.Post("/register", h.handler.Register)
	api.Post("/login", h.handler.Login)
}

func (h *httpHandler) setupProtectedRoutes(api fiber.Router) {
	api.Use(middleware.ValidateToken)
}
