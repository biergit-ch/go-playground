package controllers

import (
	"encoding/json"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/oauth"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/services"
	"git.skydevelopment.ch/zrh-dev/go-basics/config"
	"github.com/auth0/go-jwt-middleware"
	. "github.com/gorilla/mux"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Services struct {
	userService        services.UserService
	groupService       services.GroupService
	transactionService services.TransactionService
}

type Server struct {
	echo *echo.Echo
	router *Router
	services Services
	jwt *jwtmiddleware.JWTMiddleware
	conf *config.Config
}

func NewServer(userService services.UserService, groupService services.GroupService, transactionService services.TransactionService, c *config.Config) *Server {

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
		echo: nil,
	}

	return &server
}

func (api *Server) InitializeEchoHandler() *echo.Echo {

	e := echo.New()

	// Middleware
	e.Use(api.loggingMiddleware)


	api.NewUserEchoHandler(e.Group("/users"))

	return e
}


func (api *Server) InitializeHandler() *Router {

	log.Debug("Initialize Router..")

	r := NewRouter()

	// Initialize the JWT Middleware
	api.jwt = oauth.NewJwtMiddleware(api.conf)

	// inject middleware
	//r.Use(api.loggingMiddleware)
	r.Use(api.headerMiddleware)

	// initialize handler
	api.NewUserHandler(r)

	// store router in server
	api.router = r

	return api.router
}

// logging middleware for echo
func (api *Server) loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		log.WithFields(log.Fields{
			"request_uri": c.Request().RequestURI,
			"protocol": c.Request().Proto,
			"method": c.Request().Method,
		}).Debug("Calling Endpoint")
		return nil
	}
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


