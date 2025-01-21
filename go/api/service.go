package api

import (
	json2 "encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Message struct {
	Detail string `json:"detail"`
}

func (m Message) createJson(detail string) ([]byte, error) {
	m.Detail = detail
	json, err := json2.Marshal(m)
	if err != nil {
		log.Println("error while marshaling json")
		return nil, err
	}
	return json, nil
}

func MessageHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	query, err := url.ParseQuery(request.URL.RawQuery)
	log.Println("query: ", query)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("invalid query", err)
		return
	}
	detail := query.Get("message")
	if len(detail) < 2 {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("query < 2", err)
		return
	}
	var message Message
	json, err := message.createJson(detail)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while build json", err)
		return
	}
	_, err = writer.Write(json)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Println("error while write response", err)
		return
	}
}
