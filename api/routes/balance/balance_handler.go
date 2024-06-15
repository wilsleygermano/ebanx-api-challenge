package balance

import (
	cf "ebanx.api/account/domain/use_cases"
	gin "github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func BalanceHandler(c *gin.Context) {
	id := c.Query("account_id")
	balance := cf.GetBalance(id)
	if balance["error"] != nil {
		log.Println("[BalanceHandler] - Error: ", balance["error"])
		c.JSON(http.StatusNotFound, balance["ammount"])
		return
	}
	c.JSON(http.StatusOK, balance["ammount"])
}
