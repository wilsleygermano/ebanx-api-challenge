package main

import (
	b "ebanx.api/api/routes/balance"
	e "ebanx.api/api/routes/event"
	reset "ebanx.api/api/routes/reset"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/event", e.EventHandler)
	r.GET("/balance", b.BalanceHandler)
	r.POST("/reset", reset.ResetHandler)
	r.Run(":8080")
}
