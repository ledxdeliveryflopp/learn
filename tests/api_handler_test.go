package tests

import (
	"io"
	"learn/go/api"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessageHandler(t *testing.T) {
	expected := "{\"detail\":\"test\"}"
	request := httptest.NewRequest(http.MethodGet, "/test?message=test", nil)
	writer := httptest.NewRecorder()
	api.MessageHandler(writer, request)
	response := writer.Result()
	defer response.Body.Close()
	data, _ := io.ReadAll(response.Body)
	strData := string(data)
	log.Println("response status code: ", response.StatusCode)
	log.Println("response content type: ", response.Header.Get("Content-Type"))
	log.Println("response body: ", strData)
	if strData != expected {
		t.Errorf("Expected: %s, got %s", expected, string(data))
	}
}
