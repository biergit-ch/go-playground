package operations

import "fmt"

/**
links:
- https://hackernoon.com/basics-of-golang-for-beginners-6bd9b40d79ae
 */

func TestForLoop(c int) {

	persons := [3]string{"jan", "colin", "sarah"}

	for i := 0; i < len(persons); i++ {
		fmt.Println("Loop[", i, "]:", persons[i])
	}

}

func TestArray() {

	// define an array with values
	personsOne := [] string{"p1", "p2", "p3"}

	// define an array only with length
	var personsTwo [3] string
	personsTwo[0] = "jan"
	personsTwo[2] = "colin"
	personsTwo[1] = "sarah"

	fmt.Println("Array:", personsOne)
	fmt.Println("Array:", personsTwo)
}

func TestSlice(persons []string) {
	// access a specific part of an array
	var firstTwoPersons = persons[0:2]
	fmt.Println("Slice:", firstTwoPersons)
}
