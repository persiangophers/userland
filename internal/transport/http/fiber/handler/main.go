package handler

import (
	"github.com/persiangophers/userland/internal/pkg/logger"
	"github.com/persiangophers/userland/internal/service"
)

type FiberHandler struct {
	logger  logger.Logger
	account service.Account
}

func New(logger logger.Logger, accService service.Account) *FiberHandler {
	return &FiberHandler{
		logger:  logger,
		account: accService,
	}
}
