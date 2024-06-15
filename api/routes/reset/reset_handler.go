package event

import (
	"log"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func ResetHandler(c *gin.Context) {
	log.Println("[ResetHandler] - Executing")
	c.JSON(http.StatusOK, 0)
}
