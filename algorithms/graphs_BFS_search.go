package algorithms

import (
	"github.com/gammazero/deque" // либа для очередей тк в Go нет стандартного модуля
	"log"
)

func personIsSeller(person string, search string) bool {
	personRune := []rune(person)
	lastWord := string(personRune[len(personRune)-1])
	if lastWord == search {
		return true
	} else {
		return false
	}
}

func fillQueue(queue *deque.Deque[string], persons []string) {
	for _, value := range persons {
		queue.PushBack(value)
	}
}

func BFSSearchTest() bool { // поиск в ширину
	personList := make(map[string][]string) // создание хэш-сета
	searchQueue := new(deque.Deque[string]) // создание очереди
	personList["вы"] = []string{"Алиса", "Боб", "Клэр"}
	personList["Боб"] = []string{"Анудж", "Пегги"}
	personList["Алиса"] = []string{"Пегги"}
	personList["Клэр"] = []string{"Том", "Джонни"}
	personList["Анудж"] = []string{}
	personList["Пегги"] = []string{}
	personList["Том"] = []string{}
	personList["Джонни"] = []string{}
	initList := personList["вы"]     // создание 1 уровня
	fillQueue(searchQueue, initList) // добавление 1 уровня в очередь
	for searchQueue.Len() != 0 {
		person := searchQueue.PopFront()         // получение 1 объекта из очереди
		if personIsSeller(person, "и") == true { // проверка нужного условия
			log.Printf("Person %s is seller", person)
			return true
		} else {
			fillQueue(searchQueue, personList[person]) // добавление друзей объекта(2 уровень и тд)
		}
	}
	log.Println("Seller not found")
	return false
}
