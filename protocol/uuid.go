package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type UUID interface {
	Create(ctx context.Context) (entity.UUID, error)
	Delete(ctx context.Context, id entity.UUID) error
}
