package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/GersonTf/fire-backend/storage"
	"github.com/golang-jwt/jwt/v5"
)

type Server struct {
	listenAddr string
	store      storage.Storer
	jwtSecret  []byte
}

func NewServer(listenAddr string, store storage.Storer) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
		jwtSecret:  []byte("your-secret-key"), // TODO: move to config/env
	}
}

func (s *Server) Start() error {
	http.HandleFunc("GET /user/{id}", s.handleGetUserByID)
	http.HandleFunc("GET /health", s.handleHealthCheck)
	http.HandleFunc("POST /login", s.handleLogin)

	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")
	if userID == "" {
		http.Error(w, "Missing user ID in path", http.StatusBadRequest)
		return
	}

	user, err := s.store.Get(r.Context(), userID)
	if err != nil {
		// Log the error and return a 500 Internal Server Error response todo improve this strategy
		log.Printf("Failed to get user by ID: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		panic(err) //todo temporary while I define a proper strategy for errors in the handlers
	}
}

func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("I am alive!!"))
	if err != nil {
		panic(err) //todo temporary while I define a proper strategy for errors in the handlers
	}
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Create JWT token with simple claims
	claims := jwt.MapClaims{
		"user_id": "demo-user",                           // TODO: get from request body
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hour expiration
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		log.Printf("Failed to generate JWT token: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return token as JSON
	response := map[string]string{
		"token": tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err) //todo temporary while I define a proper strategy for errors in the handlers
	}
}
