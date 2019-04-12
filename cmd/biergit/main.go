package main

import (
	"context"
	"flag"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/controllers"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/services"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"git.skydevelopment.ch/zrh-dev/go-basics/playground"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// define array and initialize it with values
var persons = []string{"jan", "test1", "test2"}

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Setup Logging
	SetupLogger()

	// Test basic concepts
	BasicPrinciples()

	log.Debug("Establish DB Connection")

	// Create new MYSQL Connection
	db, err := gorm.Open("mysql", "go:123@tcp(127.0.0.1:3333)/go?charset=utf8&parseTime=True")
	db.LogMode(true)

	if err != nil {
		log.Debug("Failed to connect to database", err)
	}

	// Migrate Database
	MigrateDB(db)


	// Create Repository
	userRepo := dao.NewMysqlUserRepository(db)
	groupRepo := dao.NewMysqlGroupRepository(db)
	transactionRepo := dao.NewMysqlTransactionRepository(db)

	// Create Service
	userService := services.NewUserService(userRepo)
	groupService := services.NewGroupService(groupRepo)
	transactionService := services.NewTransactionService(transactionRepo)

	// Add some Mock Data
	jan := models.User{
		FirstName: "Jan",
		LastName:  "Minder",
	}

	// Add some Mock Data
	luca := models.User{
		FirstName: "Luca",
		LastName:  "Hostetter",
	}

	// Add some Mock Data
	biergit := models.Group{
		GroupName: "biergit",
	}

	// Add some Mock Data
	bspTrans := models.Transaction{
		Source: jan,
		Target:  luca,
		Context: biergit,
		Amount: 2,
	}

	// Create Services
	userService.CreateUser(&jan)
	groupService.CreateGroup(&biergit)
	transactionService.CreateTransaction(&bspTrans)

	// Create HTTP Server
	httpServer := controllers.NewHttpServer(userService, groupService, transactionService)
	router := httpServer.InitializeHandler()

	srv := &http.Server{
		Addr:         "0.0.0.0:8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
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
	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}

/**
Test Basic GO Principles

- Arrays
- Slices
- Loops

 */
func BasicPrinciples() {

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

func SetupLogger() {
	log.SetFormatter(&log.TextFormatter {})
	log.SetLevel(log.DebugLevel)
}

func MigrateDB(db *gorm.DB) {
	log.Debug("Migrating Database Schema")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.Transaction{})
}
