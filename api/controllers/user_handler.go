package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)


// Register User Endpoints
func (api *Server) NewUserHandler(router *mux.Router) {

	log.Debug("Initialize User Handler..")

	// Create all Users Handler
	userRouter := router.PathPrefix("/users").Subrouter()

	userRouter.Handle("/", api.GetUsers()).Methods("GET", "POST")
	userRouter.Handle("/{userId}", api.GetUser()).Methods("GET")
	userRouter.Handle("/{userId}", api.UpdateUserById()).Methods("PUT")
	userRouter.Handle("/{userId}", api.DeleteUserById()).Methods("DELETE")
}

// swagger:operation GET /users/ list
// ---
// summary: List all Users stored in repo
// description: When there are no users, it will return an empty array
func (api *Server) GetUsers() http.Handler {
	log.Debug("Initialize GET:Users Endpoint..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		users := api.services.userService.GetAllUsers()

		json.NewEncoder(w).Encode(users)
	})
}


// swagger:operation GET /users/{userId} get user by id
// ---
// summary: Get one user by its ID
// description: If the user will not be found, a 404 will be returned
// parameters:
// - name: userId
//   in: path
//   description: id of user
//   type: int
//   required: true
func (api *Server) GetUser() http.Handler {
	log.Debug("Initialize GET:UserById Endpoint..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		// Get Path Parameter
		vars := mux.Vars(r)
		userId, err := strconv.ParseInt(vars["userId"], 10, 64)

		if err != nil {
			log.Error("Failed to convert userId " , vars["userId"], " to integer")
		}

		users := api.services.userService.GetUserById(userId)

		if len(users) > 0 {
			json.NewEncoder(w).Encode(users[0])
		} else {
			http.Error(w, http.StatusText(404), 404)
		}



	})
}

func (api *Server) CreateUser() http.Handler {
	log.Debug("Initialize POST:Users Endpoint..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		users := api.services.userService.GetAllUsers()

		json.NewEncoder(w).Encode(users)
	})
}

func (api *Server) UpdateUserById() http.Handler {
	log.Debug("Initialize PUT:UserById Endpoint..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		users := api.services.userService.GetAllUsers()

		json.NewEncoder(w).Encode(users)
	})
}

// swagger:operation DELETE /users/{userId} delete user by id
// ---
// summary: Delete one user by its ID
// description: If the user will not be found, a 404 will be returned
// parameters:
// - name: userId
//   in: path
//   description: id of user
//   type: int
//   required: true
func (api *Server) DeleteUserById() http.Handler {
	log.Debug("Initialize DELETE:UserById Endpoint..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		// Get Path Parameter
		vars := mux.Vars(r)
		userId, err := strconv.ParseInt(vars["userId"], 10, 64)

		if err != nil {
			log.Error("Failed to convert userId " , vars["userId"], " to integer")
		}

		log.Debug("Delete user with id ", userId, " from repo")

		error := api.services.userService.DeleteUser(userId)

		if error != nil {
			w.WriteHeader(204)
		} else {
			w.WriteHeader(500)
		}

	})
}
