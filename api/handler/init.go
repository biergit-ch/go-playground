package handler

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/service"
	. "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Services struct {
	userService        service.UserService
	groupService       service.GroupService
	transactionService service.TransactionService
}

func NewHttpServer(userService service.UserService, groupService service.GroupService, transactionService service.TransactionService) *Services {
	return &Services{
		userService: userService,
		groupService: groupService,
		transactionService: transactionService,
	}
}

func (s *Services) InitHandler() *Router {
	log.Debug("Initialize Router..")
	r := NewRouter()
	s.NewUserHandler(r)
	s.NewGroupHandler(r)
	s.NewTransactionHandler(r)
	return r
}
