package operations

import "fmt"

func TestForLoop(c int) {

	persons := [3] string{"jan", "colin", "sarah"}

	for i := 0; i < len(persons); i++ {
		fmt.Println("Person in list[", i,"]:", persons[i])
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

	fmt.Println(personsOne)
	fmt.Println(personsTwo)
}

func TestSlice(persons [3]string) {
	// access a specific part of an array
	var firstTwoPersons = persons[0:2]
	fmt.Println(firstTwoPersons)
}
