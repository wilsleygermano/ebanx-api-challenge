package event

import (
	"log"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func ResetHandler(c *gin.Context) {
	log.Println("[ResetHandler] - Executing")
	// the answer should be
	// 200 OK
	c.String(http.StatusOK, "OK")
}
