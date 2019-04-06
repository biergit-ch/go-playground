package api

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/handler"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/service"
	. "github.com/gorilla/mux"
	"log"
)

type RestServer struct {
	userService service.UserService
	groupService service.GroupService
	transactionService service.TransactionService
}

func NewHttpServer(userService service.UserService, groupService service.GroupService, transactionService service.TransactionService) *RestServer {
	return &RestServer{
		userService: userService,
		groupService: groupService,
		transactionService: transactionService,
	}
}

func (rs *RestServer) InitHandler() *Router {

	log.Print("Initialize Router..")
	r := NewRouter()
	handler.NewUserHandler(r, rs.userService)
	handler.NewGroupHandler(r, rs.groupService)
	handler.NewTransactionHandler(r, rs.transactionService)
	return r
}
