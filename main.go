package main

import "fmt"

var persons = [3]string{"jan", "colin", "sarah"}

func main() {

	c := 15

	// test of a for loop in combination with an array
	testForLoop(c)

	testArray()

	// use of and slices
	testSlice()
}

func testForLoop(c int) {

	persons := [3] string{"jan", "colin", "sarah"}

	for i := 0; i < len(persons); i++ {
		fmt.Println("Person in list[", i,"]:", persons[i])
	}

}

func testArray() {

	personsOne := [] string{"p1", "p2", "p3"}

	var personsTwo [3] string
	personsTwo[0] = "jan"
	personsTwo[2] = "colin"
	personsTwo[1] = "sarah"

	fmt.Println(personsOne)
	fmt.Println(personsTwo)
}

func testSlice() {


	// access a specific part of an array
	var firstTwoPersons = persons[0:2]
	fmt.Println(firstTwoPersons)
}
