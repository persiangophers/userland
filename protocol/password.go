package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Password interface {
	Insert(ctx context.Context, id entity.UUID, password string) error
	Get(ctx context.Context, id entity.UUID) (string, error)
}
