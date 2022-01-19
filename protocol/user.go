package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type User interface {
	GetUser(ctx context.Context, id entity.UUID) (entity.User, error)
}
