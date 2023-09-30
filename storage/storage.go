package storage

import "github.com/GersonTf/fire-backend/types"

type Storer interface {
	Get(int) (*types.User, error)
}
