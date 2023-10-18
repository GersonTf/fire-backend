package utils

import (
	"context"
	"reflect"
	"testing"

	"github.com/GersonTf/fire-backend/config"
	"github.com/GersonTf/fire-backend/storage"
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

func CreateTestStorage(ctx context.Context, dbConn string) (storage.Storer, error) {
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

func AssertEqual(t *testing.T, want, got any) {
	t.Helper() //marking this as a testing helper is important so the test tool can report errors correctly
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func AssertNotEqual(t *testing.T, want, got any) {
	t.Helper() //marking this as a testing helper is important so the test tool can report errors correctly
	if reflect.DeepEqual(want, got) {
		t.Fatalf("Unexpected match: both values are %v", got)
	}
}
