package main

import (
	e "ebanx.api/api/routes/event"
	"github.com/gin-gonic/gin"
)

// StartServer starts the server
func main() {
	r := gin.Default()
	r.POST("/event", e.EventHandler)
	r.Run(":8080")
}
