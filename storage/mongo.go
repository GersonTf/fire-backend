package storage

import (
	"context"
	"errors"
	"time"

	"github.com/GersonTf/fire-backend/config"
	"github.com/GersonTf/fire-backend/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	mongoURI string
	client   *mongo.Client
}

func NewMongoStorage(cfg *config.Config) (*MongoStorage, error) {
	if cfg.MongoUri == "" {
		return nil, errors.New("MongoURI is empty")
	}

	clientOptions := options.Client().ApplyURI(cfg.MongoUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Optionally, you can add a ping check to ensure the connection is established.
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoStorage{
		mongoURI: cfg.MongoUri,
		client:   client,
	}, nil
}

func (s *MongoStorage) Disconnect(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}

func (s *MongoStorage) Get(id int) (*types.User, error) {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}, nil
}
