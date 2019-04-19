// GO REST API Demo
//
// This is a sample implementation with golang. This project will be used for learning's.
//
//     Schemes: http, https
//     Host: localhost:8000
//     Version: 0.1.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"context"
	"flag"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/controllers"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/repo"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/services"
	"git.skydevelopment.ch/zrh-dev/go-basics/config"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"git.skydevelopment.ch/zrh-dev/go-basics/playground"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// define array and initialize it with values
var persons = []string{"jan", "test1", "test2"}

var conf *viper.Viper

// Datasource Selection (mariadb, mongodb)
const Database = "mariadb"

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Setup Logging
	setupLogger()

	// Load configuration
	conf = config.LoadConfig("dev")

	// Test basic concepts
	basicPrinciples()

	log.Debug("Establish DB Connection")

	var userRepo repo.UserRepository
	var groupRepo repo.GroupRepository
	var transactionRepo repo.TransactionRepository

	// Check which database should be used and initiate it
	switch Database {
	case "mongodb":
		userRepo = dao.EstablishMongoDbConnection(conf)
	case "mariadb":
		userRepo, groupRepo, transactionRepo = dao.EstablishMariaDbConnection(conf)
	}

	// Create Service
	userService := services.NewUserService(userRepo)
	groupService := services.NewGroupService(groupRepo)
	transactionService := services.NewTransactionService(transactionRepo)

	// Add Mock Data
	dao.AddMockData(userService, groupService, transactionService)

	// Create HTTP Server
	httpServer := controllers.NewServer(userService, groupService, transactionService, conf)
	echoServer := httpServer.InitializeHandler()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := echoServer.Start(":" + strconv.Itoa(conf.GetInt("server.port"))); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	echoServer.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}

// Test Basic GO Principles
//
//- Arrays
//- Slices
//- Loops
func basicPrinciples() {

	log.Debug("Playground")

	// define value
	c := 15

	// test of a for loop in combination with an array
	playground.TestArray()

	playground.TestSlice(persons)

	// use of and slices
	playground.TestForLoop(c)

	// Test Function Calls
	playground.BaseCall()
	playground.WithArguments(1, 2)

	var a, b int = playground.WithMultipleReturnValues(1, 2)
	log.Debug("Multiple Return:", a, b)

	// crate an instance of Person
	user := models.User{
		FirstName: "Bier",
		LastName:  "Git",
	}

	// pass the reference of the person option
	playground.WithReferenceArguemnt(&user)

	log.Debug("Person from Main Context:", user)

	log.Debug("Playground finished")
}

func setupLogger() {
	log.SetFormatter(&log.TextFormatter {})
	log.SetLevel(log.DebugLevel)
}
