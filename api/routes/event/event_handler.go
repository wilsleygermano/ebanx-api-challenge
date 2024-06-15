package event

import (
	"log"
	"net/http"

	uc "ebanx.api/account/domain/use_cases"
	gin "github.com/gin-gonic/gin"
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

	response := uc.ChangeFunds(id, ammount)

	if response["error"] != nil {
		log.Println("[DepositHandler] - Error: ", response["error"])
		c.JSON(http.StatusNotFound, 0)
		return
	}
	c.JSON(http.StatusCreated, map[string]map[string]any{"destination": {"id": id, "balance": response["ammount"]}})
}

func WithdrawHandler(body EventRequestBody, c *gin.Context) {
	log.Println("[WithdrawHandler] - Executing")
	// POST /event {"type":"withdraw", "origin":"100", "amount":5}

	// 201 {"origin": {"id":"100", "balance":15}}

	id := body.Origin
	ammount := body.Amount

	response := uc.WithdrawFunds(id, ammount)

	if response["error"] != nil {
		log.Println("[WithdrawHandler] - Error: ", response["error"])
		c.JSON(http.StatusNotFound, 0)
		return
	}
	c.JSON(http.StatusCreated, map[string]map[string]any{"origin": {"id": id, "balance": response["ammount"]}})
}

func TransferHandler(body EventRequestBody, c *gin.Context) {
	log.Println("[TransferHandler] - Executing")
	//TODO: implement this
	c.JSON(http.StatusNotImplemented, map[string]any{"error": "not implemented"})
}
