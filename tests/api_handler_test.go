package tests

import (
	"fmt"
	"io"
	"learn/go/api"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessageHandler(t *testing.T) {
	expected := "{\"detail\":\"test\"}"
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1111/test?message=test", nil)
	log.Println(request.URL.RawQuery)
	writer := httptest.NewRecorder()
	api.MessageHandler(writer, request)
	result := writer.Result()
	defer result.Body.Close()
	data, _ := io.ReadAll(result.Body)
	fmt.Println(result.StatusCode)
	fmt.Println(result.Header.Get("Content-Type"))
	fmt.Println(string(data))
	if string(data) != expected {
		t.Errorf("Expected: %s, got %s", expected, string(data))
	}
}
