package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInsertRecord(t *testing.T) {
	// create a new gin router for testing
	router := gin.Default()

	// define the expected output of the API endpoint
	expectedResponse := "{\"code\":\"0\",\"message\":\"success\",\"primary_id\":1}\n"

	// create a request body to send to the endpoint
	requestBody := `{
        "user_id": "1234",
        "receive_amount": 100
    }`

	// create a new HTTP request with the request body
	req, err := http.NewRequest("POST", "/insert-record", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// create a new response recorder to capture the API endpoint's response
	respRecorder := httptest.NewRecorder()

	// set up a test route for the InsertRecord function
	router.POST("/insert-record", InsertRecord)

	// perform the request to the InsertRecord endpoint
	router.ServeHTTP(respRecorder, req)

	// check that the response code is what we expect
	if respRecorder.Code != http.StatusOK {
		t.Errorf("Unexpected response status code: %v", respRecorder.Code)
	}

	// check that the response body is what we expect
	if respRecorder.Body.String() != expectedResponse {
		t.Errorf("Unexpected response body: %v", respRecorder.Body.String())
	}
}
