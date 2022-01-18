package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Email interface {
	Insert(ctx context.Context, uuid [16]byte, email string) error
	Get(ctx context.Context, uuid [16]byte) (entity.Email, error)

	SendVerification(ctx context.Context, uuid [16]byte) error
	Verify(ctx context.Context, uuid [16]byte, code int) error
}
