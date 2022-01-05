package repository

import (
	"context"

	"github.com/persiangophers/userland/internal/entiry/model"
)

type Postgres interface {
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}
