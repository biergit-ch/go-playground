package controllers

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/services"
	"github.com/auth0-community/auth0"
	"github.com/auth0/go-jwt-middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/square/go-jose.v2"
	"net/http"
)

type Services struct {
	userService        services.UserService
	groupService       services.GroupService
	transactionService services.TransactionService
}

type Server struct {
	echo *echo.Echo
	services Services
	jwt *jwtmiddleware.JWTMiddleware
	conf *viper.Viper
}

func NewServer(userService services.UserService, groupService services.GroupService, transactionService services.TransactionService, c *viper.Viper) *Server {

	// Setup all Services
	services := Services{
		userService: userService,
		transactionService: transactionService,
		groupService: groupService,
	}

	// Setup Server
	server := Server{
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
	e.Use(api.headerMiddleware)

	// Initialize all sub handlers
	api.NewUserHandler(e.Group("/users"))
	api.NewTransactionHandler(e.Group("/transactions"))
	api.NewGroupHandler(e.Group("/groups"))

	return e
}

// logging middleware for echo
func (api *Server) loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.WithFields(log.Fields{
			"request_uri": c.Request().RequestURI,
			"protocol": c.Request().Proto,
			"method": c.Request().Method,
		}).Debug("Calling Endpoint")

		return next(c)
	}
}

// Use an Auth0 oAuth Provider to authorize access tokens
func (api *Server) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: api.conf.GetString("auth0.jwks")}, nil)
		log.Debug("auth0 client ", client)

		audience := api.conf.GetString("auth0.audience")
		log.Debug("auth0 audience ", audience)

		configuration := auth0.NewConfiguration(client, []string{audience}, api.conf.GetString("auth0.issuer"), jose.RS256)
		validator := auth0.NewValidator(configuration, nil)
		log.Debug("auth0 config ", configuration)

		token, err := validator.ValidateRequest(c.Request())
		log.Debug("auth0 token ", token)

		if token == nil {
			c.Error(middleware.ErrJWTMissing)
		}

		if err != nil {
			log.Debug("auth0 error ", err)
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		} else {
			log.Debug("auth0 token is valid")
			return nil
		}
	}
}

// Modify HTTP Header
func (api *Server) headerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Add("Content-Type", "application/json")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		return next(c)
	}
}

