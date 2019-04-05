package api

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/user"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/group"
	. "github.com/gorilla/mux"
)

type RestServer struct {
	userService user.UserService
	groupService group.GroupService
}

func NewHttpServer(userService user.UserService, groupService group.GroupService) *RestServer {
	return &RestServer{
		userService: userService,
		groupService: groupService,
	}
}

func (rs *RestServer) InitHandler() *Router {

	r := NewRouter()
	user.NewUserHandler(r, rs.userService)
	group.NewGroupHandler(r, rs.groupService)

	return r
}
