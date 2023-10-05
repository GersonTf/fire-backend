package api

import (
	"context"
	"encoding/json"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/GersonTf/fire-backend/config"
	"github.com/GersonTf/fire-backend/storage"
	"github.com/GersonTf/fire-backend/types"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func SetupTestMongoContainer() (*mongodb.MongoDBContainer, func(), error) {
	ctx := context.Background()

	mongodbContainer, err := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo:6"))
	if err != nil {
		return nil, nil, err
	}

	// Return cleanup function
	cleanup := func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			panic(err) // or log the error todo
		}
	}

	return mongodbContainer, cleanup, err
}

func MongoTestSetup() (*storage.MongoStorage, func(), error) {
	mc, cleanup, err := SetupTestMongoContainer()
	if err != nil {
		return nil, nil, err
	}

	cs, err := mc.ConnectionString(context.Background())
	if err != nil {
		return nil, cleanup, err
	}

	testConfig := &config.Config{
		MongoUri: cs,
		DBName:   "testdb",
	}

	store, err := storage.NewMongoStorage(testConfig)
	if err != nil {
		return nil, cleanup, err
	}

	return store, cleanup, nil
}

func TestHandleGetUsersByID(t *testing.T) {
	store, _, err := MongoTestSetup()
	if err != nil {
		t.Fatalf("Error instantiating test container: %v", err)
	}

	//defer cleanup() // Defer the cleanup function to be called after the test

	// Create test user
	testUser := &types.User{
		Name:     "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Call the Create method
	if err := store.Create(testUser); err != nil {
		t.Fatal("Error creating a test user", err)
	}

	// Create a new server using the store
	server := NewServer("", store)

	userID := testUser.ID.Hex()

	// Create a request to get the user by ID
	req, err := http.NewRequest("GET", "/user?id="+userID, nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler directly
	server.handleGetUserByID(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	// Check the response body
	var user types.User
	if err := json.Unmarshal(rr.Body.Bytes(), &user); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if user.ID.Hex() != userID {
		t.Errorf("Returned user ID does not match the inserted one: got %v want %v",
			user.ID.Hex(), userID)
	}
}
