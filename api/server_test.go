package api

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"net/http"
	"net/http/httptest"

	"github.com/GersonTf/fire-backend/storage"
	"github.com/GersonTf/fire-backend/types"
	"github.com/GersonTf/fire-backend/utils"
)

var testUser types.User
var store *storage.MongoStorage
var server *Server

func TestMain(m *testing.M) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()

	testDB, err := utils.SetupTestMongoContainer(ctx)

	if err != nil {
		panic(err)
	}

	store, err = utils.CreateTestStorage(ctx, testDB.ConStr)
	if err != nil {
		panic(err)
	}

	server = NewServer("", store)
	utils.CreateTestUser(store, &testUser)

	// Run all tests in the package
	code := m.Run()

	// I don't think we need to disconnect the db client since we are cleaning the container itself but I leave both just in case
	if cleanErr := store.Disconnect(ctx); cleanErr != nil {
		panic(cleanErr)
	}

	// Cleanup after all tests have run
	if cleanErr := testDB.Cleanup(); cleanErr != nil {
		panic(cleanErr)
	}

	// Exit with the code returned from m.Run()
	os.Exit(code)
}

func TestHandleGetUsersByID(t *testing.T) {
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

	utils.Assert(t, testUser, user)
}
