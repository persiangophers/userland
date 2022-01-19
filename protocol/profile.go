package protocol

import (
	"context"

	"github.com/persiangophers/userland/entity"
)

type Profile interface {
	CreateProfile(ctx context.Context, profile entity.Profile) error
	Get(ctx context.Context, id entity.UUID) (entity.Profile, error)
	Update(ctx context.Context, profile entity.Profile) error
}
