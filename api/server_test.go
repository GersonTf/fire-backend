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

var testUsers []*types.User //don't modify it between tests to ensure test isolation
var store storage.Storer
var server *Server

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	container, err := utils.SetupTestMongoContainer(ctx)
	if err != nil {
		panic(err)
	}

	store, err = utils.CreateTestStorage(ctx, container.ConStr)

	if err != nil {
		panic(err)
	}

	server = NewServer("", store)

	//store.Save gets the user pointer and adds the inserted db ID to it so we can use it in the tests
	testUsers = append(testUsers, &types.User{Name: "testUser", LastName: "Lawson", Email: "test@test.com", Password: "testPassword"})
	testUsers = append(testUsers, &types.User{Name: "secondU", LastName: "Foo", Email: "second@user.com", Password: "pass"})

	saveErr := store.SaveAll(ctx, testUsers)
	if saveErr != nil {
		panic(saveErr)
	}

	// Run all tests in the package
	code := m.Run()

	// I don't think we need to disconnect the db client since we are cleaning the container itself but I leave both just in case
	if cleanErr := store.Disconnect(ctx); cleanErr != nil {
		panic(cleanErr)
	}

	// Cleanup after all tests have run
	if cleanErr := container.Cleanup(); cleanErr != nil {
		panic(cleanErr)
	}

	// Exit with the code returned from m.Run()
	os.Exit(code)
}

func TestHandleGetUsersByID(t *testing.T) {
	tests := map[string]struct {
		want   types.User
		userID string
	}{
		"Getting the initial test user": {want: *testUsers[0], userID: testUsers[0].ID.Hex()},
		"Getting the second test user":  {want: *testUsers[1], userID: testUsers[1].ID.Hex()},
	}

	for tName, tt := range tests {
		t.Run(tName, func(t *testing.T) {
			// Create a request to get the user by ID
			req, err := http.NewRequest("GET", "/user?id="+tt.userID, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler directly
			server.handleGetUserByID(rr, req)

			// Check the response body
			var returnedUser types.User
			if err := json.Unmarshal(rr.Body.Bytes(), &returnedUser); err != nil {
				t.Fatalf("Failed to unmarshal response body: %v", err)
			}

			utils.AssertEqual(t, http.StatusOK, rr.Code)
			utils.AssertEqual(t, tt.userID, returnedUser.ID.Hex())
			utils.AssertEqual(t, tt.want, returnedUser)
		})
	}
}
