package handler

import (
	"encoding/json"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(router *mux.Router, s service.UserService) {
	log.Print("Initialize User Handler..")
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
