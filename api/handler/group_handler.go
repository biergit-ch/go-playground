package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Services) NewGroupHandler(router *mux.Router) {
	log.Debug("Initialize Group Handler..")
	router.Handle("/groups", s.GetGroups()).Methods("GET")
}

func (s *Services) GetGroups() http.Handler {
	log.Debug("Initialize GET:Groups Endpoint..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		groups := s.groupService.GetAllGroups()

		json.NewEncoder(w).Encode(groups)
	})
}
