package event

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	e "ebanx.api/api/routes/event"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDepositThroughEventHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/event", e.EventHandler)

	w := httptest.NewRecorder()

	event := e.EventRequestBody{
		Type:        "deposit",
		Destination: "100",
		Amount:      20,
	}

	eventJson, _ := json.Marshal(event)
	req, _ := http.NewRequest("POST", "/event", strings.NewReader(string(eventJson)))
	router.ServeHTTP(w, req)

	expectedResponse := map[string]map[string]interface{}{
		"destination": {
			"id":      "100",
			"balance": 20.0,
		},
	}

	var actualResponse map[string]map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Equal(t, expectedResponse, actualResponse)
}
