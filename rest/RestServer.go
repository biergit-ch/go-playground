package rest

import (
	"encoding/json"
	"git.skydevelopment.ch/zrh-dev/go-basics/operations"
	. "github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

/*

links:
- https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
 */

func StartRestServer(port int) {
	router := NewRouter()
	router.HandleFunc("/persons", GetPersons).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

func GetAddresses() []operations.Address {

	var addresses []operations.Address

	addresses = append(addresses, operations.Address{
		Street: "Teststrasse 4",
		City:   "City",
		Plz: 7890,
		Country: "Schweiz"})
	addresses = append(addresses, operations.Address{
		Street: "Teststrasse 4",
		City:   "City",
		Plz: 7890,
		Country: "Schweiz"})

	return addresses
}

func GetPersons(w http.ResponseWriter, r *http.Request) {

	var persons []operations.Person
	var addresses = GetAddresses()

	// create person instances
	person1 := operations.Person{
		FirstName: "Jan",
		LastName: "Minder",
		Age: 25,
		Addresses: addresses,
	}


	person2 := operations.Person{
		FirstName: "Colin",
		LastName: "Disler",
		Age:  22,
		Addresses: addresses,
	}

	persons = append(persons, person1)
	persons = append(persons, person2)


	json.NewEncoder(w).Encode(persons)
}
