package storage

import "github.com/GersonTf/fire-backend/types"

type Storer interface {
	Get(string) (*types.User, error)
	Create(*types.User) error
}
