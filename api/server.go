package api

import (
	"encoding/json"
	"net/http"

	"github.com/GersonTf/fire-backend/storage"
	"github.com/GersonTf/fire-backend/utils"
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
	//todo error? aldo get id from request
	user, _ := s.store.Get(10)

	_ = utils.Round2Dec(10.3441)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I am alive!!"))
}