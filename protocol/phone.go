package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Phone interface {
	Insert(ctx context.Context, uuid [16]byte, code, number int) error
	Get(ctx context.Context, uuid [16]byte) (entity.Phone, error)

	SendVerification(ctx context.Context, uuid [16]byte) error
	Verify(ctx context.Context, uuid [16]byte, code int) error
}
