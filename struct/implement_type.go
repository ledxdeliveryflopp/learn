package _struct

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Message struct {
	Detail string `json:"detail"`
}

func (message Message) sendJson(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(writer)
	err := jsonEncoder.Encode(message)
	if err != nil {
		message.Detail = "error"
		jsonEncoder.Encode(message)
	}
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	var message Message
	message.Detail = "test"
	message.sendJson(writer)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler).Methods("GET")
	http.ListenAndServe(":1111", r)
}
