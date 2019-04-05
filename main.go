package main

import (
	"fmt"
	"git.skydevelopment.ch/zrh-dev/go-basics/api"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/group"
	"git.skydevelopment.ch/zrh-dev/go-basics/operations"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/user"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

// define array and initialize it with values
var persons = []string{"jan", "test1", "test2"}

func main() {

	// Test basic concepts
	BasicPrinciples()

	// Create new MYSQL Connection
	db, err := gorm.Open("mysql", "go:123@tcp(127.0.0.1:3333)/go-basics?charset=utf8&parseTime=True")

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	// Migrate Database
	MigrateDB(db)


	// Create Repository
	userRepo := user.NewMysqlRepository(db)
	groupRepo := group.NewMysqlRepository(db)

	// Create Service
	userService := user.NewUserService(userRepo)
	groupService := group.NewGroupService(groupRepo)

	// Add some Mock Data
	jan := user.User{
		FirstName: "Jan",
		LastName:  "Minder",
	}

	// Add some Mock Data
	biergit := group.Group{
		GroupName: "biergit",
	}

	// Create New User
	userService.CreateUser(&jan)
	groupService.CreateGroup(&biergit)

	httpServer := api.NewHttpServer(userService, groupService)
	router := httpServer.InitHandler()

	http.ListenAndServe("127.0.0.1:8000", router )
}

/**
Test Basic GO Principles

- Arrays
- Slices
- Loops

 */
func BasicPrinciples() {

	// define value
	c := 15

	// test of a for loop in combination with an array
	operations.TestArray()

	operations.TestSlice(persons)

	// use of and slices
	operations.TestForLoop(c)

	// Test Function Calls
	operations.BaseCall()
	operations.WithArguments(1, 2)

	var a, b int = operations.WithMultipleReturnValues(1, 2)
	fmt.Println("Multiple Return:", a, b)

	// crate an instance of Person
	user := user.User{
		FirstName: "Bier",
		LastName:  "Git",
	}

	// pass the reference of the person option
	operations.WithReferenceArguemnt(&user)

	fmt.Println("Person from Main Context:", user)
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&group.Group{})
}
