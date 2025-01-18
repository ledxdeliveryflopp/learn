package learn_struct

// Модуль для изучения структур

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Message struct {
	Detail string `json:"detail"`
} // Что бы избежать фрагментации/выравнивания(выделение лишней памяти), нужно, что бы самые маленькие(по байтам)
// поля были в конце

// Добавляем метод структуре, что бы в каждом эндпоинте api не вызывать метод для установки Content-Type и инициализации json encoder
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
	message.sendJson(writer) // используем метод структуры
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler).Methods("GET")
	http.ListenAndServe(":1111", r)
}
