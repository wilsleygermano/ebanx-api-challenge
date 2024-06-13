package event

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ebanx.api/api/routes/event"
	e "ebanx.api/api/routes/event"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// func TestPostUser(t *testing.T) {
// 	router := setupRouter()
// 	router = postUser(router)

// 	w := httptest.NewRecorder()

// 	// Create an example user for testing
// 	exampleUser := User{
// 		Username: "test_name",
// 		Gender:   "male",
// 	}
// 	userJson, _ := json.Marshal(exampleUser)
// 	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	// Compare the response body with the json data of exampleUser
// 	assert.Equal(t, string(userJson), w.Body.String())
// }

func TestDepositThroughEventHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/event", e.EventHandler)

	w := httptest.NewRecorder()

	event := event.EventRequestBody{
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
			"balance": 20.0, // json.Unmarshal converts numbers to float64
		},
	}

	var actualResponse map[string]map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Equal(t, expectedResponse, actualResponse)
}
