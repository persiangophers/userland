package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Phone interface {
	Insert(ctx context.Context, id entity.UUID, code, number int) error
	Get(ctx context.Context, id entity.UUID) (entity.Phone, error)
}
