package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Password interface {
	Insert(ctx context.Context, uuid entity.UUID, password string) error
	Update(ctx context.Context, uuid entity.UUID, password string) error

	Verify(ctx context.Context, uuid entity.UUID, password string) (bool, error)
}
