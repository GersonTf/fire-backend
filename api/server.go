package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GersonTf/fire-backend/storage"
)

type Server struct {
	listenAddr string
	store      storage.Storer
}

func NewServer(listenAddr string, store storage.Storer) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserByID)
	http.HandleFunc("/health", s.handleHealthCheck)

	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	//todo errors? aldo get id from request
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	user, err := s.store.Get(userID)
	if err != nil {
		// Log the error and return a 500 Internal Server Error response todo improve this strategy
		log.Printf("Failed to get user by ID: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I am alive!!"))
}
