package operations

import "fmt"

/*

Passing Reference or a copy of the variable ?

Even though Go looks a bit like C, its compiler works differently. And C
analogy does not always work with Go. Passing by value in Go may be significantly
cheaper than passing by pointer. This happens because Go uses escape analysis to determine
if variable can be safely allocated on function’s stack frame, which could be
much cheaper then allocating variable on the heap. Passing by value simplifies
escape analysis in Go and gives variable a better chance to be allocated on the stack.

src: http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/

*/


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
	person.FirstName = "manipulated from function jan"
}
