package operations

import "fmt"

func BaseCall() int {
	return 1
}

func WithArguments(a int, b int) int  {
 return a + b
}


/**
method with multiple return commands
  */
func WithMultipleReturnValues(a int, b int) (int, int) {
	return a * 5, b * 4
}

func WithReferenceArguemnt(person *Person) {
	fmt.Println("Print Person Value from reference", person)
}
