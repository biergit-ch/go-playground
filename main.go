package main

import (
	"fmt"
	"git.skydevelopment.ch/zrh-dev/go-basics/api"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/model"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/service"
	"git.skydevelopment.ch/zrh-dev/go-basics/playground"
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
	userRepo := dao.NewMysqlUserRepository(db)
	groupRepo := dao.NewMysqlGroupRepository(db)
	transactionRepo := dao.NewMysqlTransactionRepository(db)

	// Create Service
	userService := service.NewUserService(userRepo)
	groupService := service.NewGroupService(groupRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	// Add some Mock Data
	jan := model.User{
		FirstName: "Jan",
		LastName:  "Minder",
	}

	// Add some Mock Data
	luca := model.User{
		FirstName: "Luca",
		LastName:  "Hostetter",
	}

	// Add some Mock Data
	biergit := model.Group{
		GroupName: "biergit",
	}

	// Add some Mock Data
	bspTrans := model.Transaction{
		Source: jan,
		Target:  luca,
		Context: biergit,
		Amount: 2,
	}

	// Create New User
	userService.CreateUser(&jan)
	groupService.CreateGroup(&biergit)
	transactionService.CreateTransaction(&bspTrans)

	httpServer := api.NewHttpServer(userService, groupService, transactionService)
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

	log.Println("Playground")

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
	fmt.Println("Multiple Return:", a, b)

	// crate an instance of Person
	user := model.User{
		FirstName: "Bier",
		LastName:  "Git",
	}

	// pass the reference of the person option
	playground.WithReferenceArguemnt(&user)

	fmt.Println("Person from Main Context:", user)

	log.Println("Playground finished")
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Group{})
	db.AutoMigrate(&model.Transaction{})
}
