package controllers

import (
	"encoding/json"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/oauth"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/services"
	"git.skydevelopment.ch/zrh-dev/go-basics/config"
	"github.com/auth0/go-jwt-middleware"
	. "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Services struct {
	userService        services.UserService
	groupService       services.GroupService
	transactionService services.TransactionService
}

type Server struct {
	router *Router
	services Services
	jwt *jwtmiddleware.JWTMiddleware
	conf *config.Config
}

func NewHttpServer(userService services.UserService, groupService services.GroupService, transactionService services.TransactionService, c *config.Config) *Server {

	// Setup all Services
	services := Services{
		userService: userService,
		transactionService: transactionService,
		groupService: groupService,
	}

	// Setup Server
	server := Server{
		router: nil,
		services: services,
		conf: c,
	}

	return &server
}


func (api *Server) InitializeHandler() *Router {

	log.Debug("Initialize Router..")

	r := NewRouter()

	// Initialize the JWT Middleware
	api.jwt = oauth.NewJwtMiddleware(api.conf)

	// inject middleware
	r.Use(api.loggingMiddleware)
	r.Use(api.headerMiddleware)

	// initialize handler
	api.NewUserHandler(r)

	// store router in server
	api.router = r

	return api.router
}

// Logging Middleware for all HTTP Requests
func (api *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"request_uri": r.RequestURI,
			"protocol": r.Proto,
			"method": r.Method,
		}).Debug("Calling Endpoint")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// Modify HTTP Header
func (api *Server) headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "application/json")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func (api *Server) responseJSON(message string, w http.ResponseWriter, statusCode int) {
	response := oauth.Response{message}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}


