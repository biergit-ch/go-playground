package main

import (
	"fmt"
	"git.skydevelopment.ch/zrh-dev/go-basics/http"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"git.skydevelopment.ch/zrh-dev/go-basics/operations"
	"git.skydevelopment.ch/zrh-dev/go-basics/dao"
)

// define array and initialize it with values
var persons = []string{"jan", "test1", "test2"}

func main() {

	// Test basic concepts
	BasicPrinciples()

	// crate an instance of Person
	user := models.User{
		FirstName: "Bier",
		LastName:  "Git",
	}

	// pass the reference of the person option
	operations.WithReferenceArguemnt(&user)

	fmt.Println("Person from Main Context:", user)

	// Setup a mysql database connection
	SetupDatabase()

	// Start REST Webserver
	http.StartRestServer(8000)
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
}

//TODO: Close Database Connection!
func SetupDatabase() {

	// Initialize Database
	dao.InitDB()

	// Add some Mock Data
	dao.AddMockData()
}


