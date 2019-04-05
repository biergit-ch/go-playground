package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	Service UserService
}

func NewUserHandler(router *mux.Router, s UserService) {
	handler := &UserHandler{
		Service: s,
	}
	router.Handle("/users", GetUsers(handler)).Methods("GET")
}

func GetUsers(env *UserHandler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		users := env.Service.GetAllUsers()

		json.NewEncoder(w).Encode(users)
	})
}
