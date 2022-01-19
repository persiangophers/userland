package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Phone interface {
	Insert(ctx context.Context, uuid entity.UUID, code, number int) error
	Get(ctx context.Context, uuid entity.UUID) (entity.Phone, error)

	SendVerification(ctx context.Context, uuid entity.UUID) error
	Verify(ctx context.Context, uuid entity.UUID, code int) error
}
