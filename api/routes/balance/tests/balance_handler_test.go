package event

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	b "ebanx.api/api/routes/balance"
	e "ebanx.api/api/routes/event"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountBalanceThroughHandler(t *testing.T) {
	router := gin.Default()
	router.GET("/balance", b.BalanceHandler)
	router.POST("/event", e.EventHandler)

	// POST request
	wPost := httptest.NewRecorder()
	event := e.EventRequestBody{
		Type:        "deposit",
		Destination: "123",
		Amount:      20,
	}
	eventJson, _ := json.Marshal(event)
	reqPost, _ := http.NewRequest("POST", "/event", strings.NewReader(string(eventJson)))
	router.ServeHTTP(wPost, reqPost)
	assert.Equal(t, http.StatusCreated, wPost.Code)

	// GET request
	wGet := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/balance?account_id=123", nil)
	router.ServeHTTP(wGet, req)
	assert.Equal(t, http.StatusOK, wGet.Code)

}
