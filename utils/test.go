package utils

import (
	"context"
	"reflect"
	"testing"

	"github.com/GersonTf/fire-backend/config"
	"github.com/GersonTf/fire-backend/storage"
	"github.com/GersonTf/fire-backend/types"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

const TestDBName string = "testDB"

// testDB containes the dockerized mongodb,
// the cleanup function to clean the container after use and the connection string
type TestDB struct {
	ConStr    string                    //generated db container connection String
	Container *mongodb.MongoDBContainer //the container itself in which the DB is running
	Cleanup   func() error              //cleanup function to clean the container after use
}

func SetupTestMongoContainer(ctx context.Context) (*TestDB, error) {
	mongodbContainer, initErr := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo:6"))
	if initErr != nil {
		return nil, initErr
	}

	// Return cleanup function
	cleanup := func() error {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			return (err)
		}
		return nil
	}

	cs, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}

	return &TestDB{ConStr: cs, Container: mongodbContainer, Cleanup: cleanup}, err
}

func CreateTestStorage(ctx context.Context, dbConn string) (*storage.MongoStorage, error) {
	testConfig := &config.Config{
		MongoUri: dbConn,
		DBName:   TestDBName,
	}

	store, err := storage.NewMongoStorage(ctx, testConfig)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func CreateTestUser(store *storage.MongoStorage, user *types.User) error {
	// Create test user
	user.Name = "testuser"
	user.Email = "test@example.com"
	user.Password = "password123"

	//create gets the user pointer and adds the inserted db ID to it
	if err := store.Create(context.Background(), user); err != nil {
		return err
	}

	return nil
}

func Assert(t *testing.T, want any, got any) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func AssertNotEqual(t *testing.T, want any, got any) {
	if reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected match: both values are %v", got)
	}
}
