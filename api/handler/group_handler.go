package handler

import (
	"encoding/json"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type GroupHandler struct {
	Service service.GroupService
}

func NewGroupHandler(router *mux.Router, s service.GroupService) {
	log.Print("Initialize Group Handler..")
	handler := &GroupHandler{
		Service: s,
	}
	router.Handle("/groups", GetGroups(handler)).Methods("GET")
}

func GetGroups(env *GroupHandler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		groups := env.Service.GetAllGroups()

		json.NewEncoder(w).Encode(groups)
	})
}
