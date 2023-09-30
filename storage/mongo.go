package storage

import "github.com/GersonTf/fire-backend/types"

type MongoStorage struct{}

func (s *MongoStorage) Get(id int) (*types.User, error) {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}, nil
}
