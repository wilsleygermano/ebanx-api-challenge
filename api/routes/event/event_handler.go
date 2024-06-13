package event

import (
	cf "ebanx.api/account/domain/use_cases"
	gin "github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EventHandler(c *gin.Context) {
	var body EventRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		log.Println("[EventHandler] - Error: ", err)
		c.JSON(http.StatusBadRequest, 0)
	}
	event := body.Type
	switch event {
	case "deposit":
		DepositHandler(body, c)
	case "withdraw":
		WithdrawHandler(body, c)
	case "transfer":
		TransferHandler(body, c)
	default:
		c.JSON(http.StatusBadRequest, 0)

	}
}

func DepositHandler(body EventRequestBody, c *gin.Context) {
	log.Println("[DepositHandler] - Executing")
	id := body.Destination
	ammount := body.Amount

	err := cf.ChangeFunds(id, ammount)
	if err != nil {
		log.Println("[DepositHandler] - Error: ", err)
		c.JSON(http.StatusBadRequest, 0)
	}
	c.JSON(http.StatusCreated, map[string]map[string]any{"destination": {"id": id, "balance": ammount}})
}

func WithdrawHandler(body EventRequestBody, c *gin.Context) {
	log.Println("[WithdrawHandler] - Executing")
	// TODO: implement this
	c.JSON(http.StatusNotImplemented, map[string]any{"error": "not implemented"})
}

func TransferHandler(body EventRequestBody, c *gin.Context) {
	log.Println("[TransferHandler] - Executing")
	//TODO: implement this
	c.JSON(http.StatusNotImplemented, map[string]any{"error": "not implemented"})
}
