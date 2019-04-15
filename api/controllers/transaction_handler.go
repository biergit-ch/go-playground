package controllers

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (api *Server) NewTransactionHandler(g *echo.Group) {
	log.Debug("Initialize Transaction Handler..")

	g.Use(api.authMiddleware)
	g.GET("/private", api.GetTransactions)
	g.GET("/:id", api.GetTransaction)
	log.Debug("Transaction Handler initialized ", g)
}

// swagger:operation GET /transactions list
// ---
// summary: List all Transactions stored in repo
// description: When there are no transactions, it will return an empty array
func (api *Server) GetTransactions(c echo.Context) error {
	log.Debug("Get Transactions")
	transactions := api.services.transactionService.GetAllTransactions()
	return c.JSON(http.StatusOK ,transactions)
}

// swagger:operation GET /transactions/{transactionId} get one transaction
// ---
// summary: List all Transactions stored in repo
// description: When there are no groups, it will return an empty array
// parameters:
// - name: transactionId
//   in: path
//   description: transaction id
//   type: int
//   required: true
func (api *Server) GetTransaction(c echo.Context) error {
	log.Debug("Get Transaction")

	id, err := strconv.Atoi(c.Param("id"))
	log.Debug("User id is ", id)

	if err != nil {
		log.Error("Failed to convert userId " , c.Param("id"), " to int64")
	}

	//transaction := api.services.transactionService.GetTransaction(id)
	return c.JSON(http.StatusOK ,nil)
}
