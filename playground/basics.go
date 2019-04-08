package playground

import (
	log "github.com/sirupsen/logrus"
)

func TestForLoop(c int) {

	persons := [3]string{"jan", "test1", "test2"}

	for i := 0; i < len(persons); i++ {
		log.Println("Loop[", i, "]:", persons[i])
	}

}

func TestArray() {

	// define an array with values
	personsOne := [] string{"p1", "p2", "p3"}

	// define an array only with length
	var personsTwo [3] string
	personsTwo[0] = "jan"
	personsTwo[2] = "test1"
	personsTwo[1] = "test2"

	log.Println("Array:", personsOne)
	log.Println("Array:", personsTwo)
}

func TestSlice(persons []string) {
	// access a specific part of an array
	var firstTwoPersons = persons[0:2]
	log.Println("Slice:", firstTwoPersons)
}
