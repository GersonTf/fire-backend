package storage

import (
	"github.com/GersonTf/fire-backend/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (ms *MemoryStorage) Get(id string) (*types.User, error) {
	return &types.User{
		ID:   primitive.NewObjectID(),
		Name: "Foo",
	}, nil
}
