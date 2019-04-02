package main

import (
	"fmt"
	"git.skydevelopment.ch/zrh-dev/go-basics/operations"
	"git.skydevelopment.ch/zrh-dev/go-basics/rest"
)

// define array and initialize it with values
var persons = []string{"jan", "colin", "sarah"}

func main() {

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
	person := operations.Person {
		FirstName: "Jan",
		LastName: "Minder",
		Age: 25,
	}

	// pass the reference of the person option
	operations.WithReferenceArguemnt(&person)

	fmt.Println("Person from Main Context:", person)

	// Start REST Webserver
	rest.StartRestServer(8000)
}

