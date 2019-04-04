package controller

import (
	"encoding/json"
	"git.skydevelopment.ch/zrh-dev/go-basics/dao"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	var user = dao.GetUser()

	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dao.GetUsers())
}
