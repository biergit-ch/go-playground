package handler

import (
	"encoding/json"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TransactionHandler struct {
	Service service.TransactionService
}

func NewTransactionHandler(router *mux.Router, s service.TransactionService) {
	log.Print("Initialize Transaction Handler..")
	handler := &TransactionHandler{
		Service: s,
	}
	router.Handle("/transactions", GetTransactions(handler)).Methods("GET")
}

func GetTransactions(env *TransactionHandler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		transactions := env.Service.GetAllTransactions()

		json.NewEncoder(w).Encode(transactions)
	})
}
