package service

import (
	"context"

	"github.com/persiangophers/userland/internal/entiry/model"
)

type Account interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, user *model.User) error
}
