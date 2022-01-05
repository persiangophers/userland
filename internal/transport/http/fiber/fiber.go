package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/persiangophers/userland/internal/pkg/logger"
	"github.com/persiangophers/userland/internal/service"
	"github.com/persiangophers/userland/internal/transport/http"
	"github.com/persiangophers/userland/internal/transport/http/fiber/handler"
)

type httpHandler struct {
	app     *fiber.App
	handler *handler.FiberHandler
}

func New(logger logger.Logger, accService service.Account) http.HttpHandler {
	app := fiber.New()
	return &httpHandler{
		app:     app,
		handler: handler.New(logger, accService),
	}
}

func (h *httpHandler) Start(address string) error {
	h.setupRoutes(h.app)
	h.app.Listen(address)
	return nil
}

func (h *httpHandler) Shutdown() error {
	return h.app.Shutdown()
}
