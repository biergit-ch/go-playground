package rest

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (api *Server) NewGroupHandler(g *echo.Group) {
	log.Debug("Initialize Group Handler..")

	g.GET("", api.GetGroups)
}

// swagger:operation GET /groups list
// ---
// summary: List all Groups stored in repo
// description: When there are no groups, it will return an empty array
func  (api *Server) GetGroups(c echo.Context) error {
	log.Debug("Get all Groups")

	groups := api.services.groupService.GetAllGroups()
	return c.JSON(http.StatusOK ,groups)
}
