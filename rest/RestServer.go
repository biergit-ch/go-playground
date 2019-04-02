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
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), router))
}

func GetPersons(w http.ResponseWriter, r *http.Request) {

	var persons []operations.Person

	persons = append(persons, operations.Person{"Jan", "Minder", 25})
	persons = append(persons, operations.Person{"Colin", "Disler", 22})
	json.NewEncoder(w).Encode(persons)
}
