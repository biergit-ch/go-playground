package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Services) NewUserHandler(router *mux.Router) {
	log.Debug("Initialize User Handler..")
	router.Handle("/transactions", s.GetUsers()).Methods("GET")
}

func (s *Services) GetUsers() http.Handler {
	log.Debug("Initialize GET:Users Endpoint..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		users := s.userService.GetAllUsers()

		json.NewEncoder(w).Encode(users)
	})
}
