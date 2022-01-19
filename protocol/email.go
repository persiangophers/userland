package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Email interface {
	Insert(ctx context.Context, id entity.UUID, email string) error
	Get(ctx context.Context, id entity.UUID) (entity.Email, error)
}
