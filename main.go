package main

import (
	b "ebanx.api/api/routes/balance"
	e "ebanx.api/api/routes/event"
	"github.com/gin-gonic/gin"
)

// StartServer starts the server
func main() {
	r := gin.Default()
	r.POST("/event", e.EventHandler)
	r.GET("/balance", b.BalanceHandler)
	r.Run(":8080")
}
