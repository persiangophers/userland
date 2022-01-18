package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type User interface {
	GetUser(ctx context.Context, uuid [16]byte) (entity.User, error)
}
