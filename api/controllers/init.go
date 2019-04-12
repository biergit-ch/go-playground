package controllers

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/services"
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
}

func NewHttpServer(userService services.UserService, groupService services.GroupService, transactionService services.TransactionService) *Server {

	// Setup all Services
	services := Services{
		userService: userService,
		transactionService: transactionService,
		groupService: groupService,
	}

	// Setup Server
	server := Server{
		router: NewRouter(),
		services: services,
	}

	return &server
}


func (api *Server) InitializeHandler() *Router {

	log.Debug("Initialize Router..")

	r := NewRouter()

	// inject middleware
	r.Use(api.loggingMiddleware)

	// initialize handler
	api.NewUserHandler(r)

	return r
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
