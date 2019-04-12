package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)


func (s *Services) NewTransactionHandler(router *mux.Router) {
	log.Debug("Initialize Transaction Handler..")
	router.Handle("/transactions", s.GetTransactions()).Methods("GET")
	router.Handle("/transactions", s.GetTransactions()).Methods("GET")
}

func (s *Services) GetTransactions() http.Handler {
	log.Debug("Initialize GET:Transaction Endpoint..")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}

		transactions := s.transactionService.GetAllTransactions()

		json.NewEncoder(w).Encode(transactions)
	})
}
