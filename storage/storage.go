package storage

import (
	"context"

	"github.com/GersonTf/fire-backend/types"
)

type Storer interface {
	Get(context.Context, string) (*types.User, error)
	Save(context.Context, *types.User) error
	Disconnect(ctx context.Context) error
}
