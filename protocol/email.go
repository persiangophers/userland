package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Email interface {
	Insert(ctx context.Context, uuid entity.UUID, email string) error
	Get(ctx context.Context, uuid entity.UUID) (entity.Email, error)

	SendVerification(ctx context.Context, uuid entity.UUID) error
	Verify(ctx context.Context, uuid entity.UUID, code int) error
}
