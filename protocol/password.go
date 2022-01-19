package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Password interface {
	Insert(ctx context.Context, id entity.UUID, password string) error
	Update(ctx context.Context, id entity.UUID, password string) error

	Verify(ctx context.Context, id entity.UUID, password string) (bool, error)
}
