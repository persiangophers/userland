package protocol

import "context"

type Password interface {
	Insert(ctx context.Context, uuid [16]byte, password string) error
	Update(ctx context.Context, uuid [16]byte, password string) error

	Verify(ctx context.Context, uuid [16]byte, password string) (bool, error)
}
