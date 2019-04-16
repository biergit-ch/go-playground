package controllers

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (api *Server) NewUserHandler(g *echo.Group) {
	g.GET("", api.GetUsers)
	g.POST("", api.CreateUser)
	g.GET("/:id", api.GetUser)
	g.DELETE("/:id", api.DeleteUser)
}

// swagger:operation GET /users list
// ---
// summary: List all Users stored in repo
// description: When there are no users, it will return an empty array
func (api *Server) GetUsers(c echo.Context) error {
	users := api.services.userService.GetAllUsers()
	return c.JSON(http.StatusOK, users)
}


// swagger:operation GET /users/{userId} get user by id
// ---
// summary: Get one user by its ID
// description: If the user will not be found, a 404 will be returned
// parameters:
// - name: userId
//   in: path
//   description: id of user
//   type: int
//   required: true
func (api *Server) GetUser(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("id"))

	users := api.services.userService.GetUserById(userId)

	if len(users) > 0 {
		return c.JSON(http.StatusOK, users[0])
	} else {
		return c.NoContent(http.StatusNotFound)
	}
}


// swagger:operation POST /users create new user
// ---
// summary: Create a new User
func (api *Server) CreateUser(c echo.Context) error {

	u := &models.User{}

	if err := c.Bind(u); err != nil {
		return err
	}

	api.services.userService.CreateUser(u)

	return c.JSON(http.StatusCreated, u)

}


// swagger:operation PUT /users/{userId} update user by id
// ---
// summary: Delete one user by its ID
// description: If the user will not be found, a 404 will be returned
// parameters:
// - name: userId
//   in: path
//   description: id of user
//   type: int
//   required: true
func (api *Server) UpdateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	userId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	return c.JSON(http.StatusOK, userId)
}

// swagger:operation DELETE /users/{userId} delete user by id
// ---
// summary: Delete one user by its ID
// description: If the user will not be found, a 404 will be returned
// parameters:
// - name: userId
//   in: path
//   description: id of user
//   type: int
//   required: true
func  (api *Server) DeleteUser(c echo.Context) error {

	id, convErr := strconv.Atoi(c.Param("id"))

	if convErr != nil {
		log.Error("Failed to convert userId " , c.Param("id"), " to int64")
	}

	log.Debug("Try to delete user with id ", id, " from repo")

	delErr := api.services.userService.DeleteUser(id)

	if delErr != nil {
		log.Error("Failed to delete user ", delErr)
		return c.NoContent(http.StatusInternalServerError)
	} else {
		return c.NoContent(http.StatusNoContent)
	}
}
