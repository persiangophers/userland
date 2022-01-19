package protocol

import "context"

type UUID interface {
	Create(ctx context.Context) ([16]byte, error)
	Delete(ctx context.Context, uuid []byte) error
}
