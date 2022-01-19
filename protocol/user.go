package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type User interface {
	GetUser(ctx context.Context, uuid entity.UUID) (entity.User, error)
}
