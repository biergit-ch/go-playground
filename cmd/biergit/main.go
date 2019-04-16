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
	"git.skydevelopment.ch/zrh-dev/go-basics/api/services"
	"git.skydevelopment.ch/zrh-dev/go-basics/config"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"git.skydevelopment.ch/zrh-dev/go-basics/playground"
	"github.com/jinzhu/gorm"
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

func main() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	// Setup Logging
	SetupLogger()

	// Load configuration
	conf = config.LoadConfig("dev")

	// Test basic concepts
	BasicPrinciples()

	log.Debug("Establish DB Connection")

	// Create new MYSQL Connection
	db, err := gorm.Open("mysql", conf.GetString("database.user") + ":" + conf.GetString("database.password") + "@tcp(" + conf.GetString("database.host") + ":" + strconv.Itoa(conf.GetInt("database.port")) + ")/" + conf.GetString("database.schema") + "?charset=utf8&parseTime=True")

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// Enable SQL Query Logs
	db.LogMode(true)

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

	// Add Mock Data
	addMockData(userService, groupService, transactionService)

	// Create HTTP Server
	httpServer := controllers.NewServer(userService, groupService, transactionService, conf)
	echoServer := httpServer.InitializeHandler()

	//echoServer.Logger.Fatal(echoServer.Start(":8001"))

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


func addMockData(userService services.UserService, groupService services.GroupService, transactionService services.TransactionService) {

	log.Debug("Add Mock data")

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
