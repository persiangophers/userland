package http

import "github.com/persiangophers/userland/internal/config"

type IRest interface {
	Start(address string, cfg *config.Config) error
	Shutdown() error
}
