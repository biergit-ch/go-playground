package group

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type GroupHandler struct {
	Service GroupService
}

func NewGroupHandler(router *mux.Router, s GroupService) {
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
