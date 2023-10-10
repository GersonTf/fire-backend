package storage

import (
	"context"

	"github.com/GersonTf/fire-backend/types"
)

type Storer interface {
	Get(context.Context, string) (*types.User, error)
	Create(context.Context, *types.User) error
}
