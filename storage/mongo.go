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

func NewMongoStorage(parentCtx context.Context, cfg *config.Config) (*MongoStorage, error) {
	if cfg.MongoUri == "" || cfg.DBName == "" {
		return nil, errors.New("missing DB config variables")
	}

	clientOptions := options.Client().ApplyURI(cfg.MongoUri)
	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
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

func (s *MongoStorage) Get(ctx context.Context, id string) (*types.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User

	collection := s.db.Collection(UserCollection)
	filter := bson.M{"_id": objID}

	err = collection.FindOne(ctx, filter).Decode(&user)

	switch err {
	case mongo.ErrNoDocuments:
		return nil, fmt.Errorf("user %s not found", id)
	case nil:
		return &user, nil
	default:
		return nil, err
	}

}

// todo should be in user
func (s *MongoStorage) Save(ctx context.Context, user *types.User) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	collection := s.db.Collection(UserCollection) //todo this could probably live outside of the func

	result, err := collection.InsertOne(ctx, bson.M{
		"username": user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		return err
	}

	// Update the ID field of the user argument
	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// saveAll assumes the users slice don't have generated IDs
func (s *MongoStorage) SaveAll(ctx context.Context, users []*types.User) error {
	documents := make([]interface{}, len(users))

	for i, user := range users {
		// Generate new ObjectID for the user
		user.ID = primitive.NewObjectID()

		documents[i] = bson.M{
			"_id":      user.ID,
			"name":     user.Name,
			"lastName": user.LastName,
			"email":    user.Email,
			"password": user.Password,
		}
	}

	collection := s.db.Collection(UserCollection) //todo this could probably live outside of the func
	_, err := collection.InsertMany(ctx, documents)
	return err
}
