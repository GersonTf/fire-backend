package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GersonTf/fire-backend/config"
	"github.com/GersonTf/fire-backend/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const UserCollection string = "user"

type MongoStorage struct {
	mongoURI string
	client   *mongo.Client
	db       *mongo.Database
}

func NewMongoStorage(cfg *config.Config) (*MongoStorage, error) {
	if cfg.MongoUri == "" || cfg.DBName == "" {
		return nil, errors.New("missing DB config variables")
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
		db:       client.Database(cfg.DBName),
	}, nil
}

func (s *MongoStorage) Disconnect(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}

func (s *MongoStorage) Get(id string) (*types.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User

	collection := s.db.Collection(UserCollection)

	filter := bson.M{"_id": objID}

	//todo learn about context
	err = collection.FindOne(context.TODO(), filter).Decode(&user)

	switch err {
	case mongo.ErrNoDocuments:
		return nil, fmt.Errorf("User %s not found", id)
	case nil:
		// No error, so do nothing
	default:
		return nil, err
	}

	return &user, nil
}
