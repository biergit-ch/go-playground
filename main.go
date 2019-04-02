package main

import "git.skydevelopment.ch/zrh-dev/go-basics/operations"

var persons = [3]string{"jan", "colin", "sarah"}

func main() {

	c := 15

	// test of a for loop in combination with an array
	operations.TestArray()

	operations.TestSlice(persons)

	// use of and slices
	operations.TestForLoop(c)

	// Test Function Calls
	operations.BaseCall()
	operations.WithArguments(1, 2)

}

