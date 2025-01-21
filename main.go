package main

import (
	"learn/go/api"
	"log"
	"math/rand"
	"net/http"
)

// Создание массива определенной длины со случайными цифрами в определенном диапазоне
func createArray(listLen int, valueRange int) []int {
	myArray := make([]int, listLen)
	for i := range myArray {
		myArray[i] = rand.Intn(valueRange)
	}
	return myArray
}

func main() {
	//arr := createArray(100, 241) // создание массива
	//log.Println("new list: ", arr)
	//log.Println("new list len: ", len(arr))
	//sortedArr := algorithms.QuickSort([]int{1, 4, 5, 12, 3, 2, 6, 1, 4, 13}) // сортировка массива
	//log.Println("========================================")
	//log.Println("sorted list: ", sortedArr)
	//log.Println("sorted list len: ", len(sortedArr))
	//index, value, status := algorithms.BinarySearch(sortedArr, 21) // бинарный поиск
	//if status == true {
	//	log.Printf("index: %d, value: %d ", index, value)
	//	log.Printf("check, get value by index: %d", sortedArr[index])
	//}
	//algorithms.BFSSearchTest()
	http.HandleFunc("/test", api.MessageHandler)
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

//func main() {
//	result := recursion.FirstRecursionFactorial(5)
//	log.Println(result)
//}
