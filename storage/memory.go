package storage

import "github.com/GersonTf/fire-backend/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (ms *MemoryStorage) Get(id int) (*types.User, error) {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}, nil
}
