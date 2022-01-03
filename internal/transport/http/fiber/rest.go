package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/transport/http"
	"go.uber.org/zap"
)

type rest struct {
	fiber  *fiber.App
	logger *zap.Logger
}

func New(logger *zap.Logger) http.IRest {
	return &rest{
		fiber:  fiber.New(),
		logger: logger,
	}
}

func (r *rest) Start(address string, cfg *config.Config) error {
	r.routing(cfg)
	return r.fiber.Listen(address)
}

func (r *rest) Shutdown() error {
	//TODO implement me
	panic("implement me")
}
