package http

import (
	"encoding/json"
	"git.skydevelopment.ch/zrh-dev/go-basics/http/controller"
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
	r := NewRouter()
	r.HandleFunc("/", GetIndex).Methods("GET")
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Go Index Location")
}

