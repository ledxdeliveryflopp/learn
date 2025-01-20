package main

import (
	"learn/algorithms"
	"log"
	"math/rand"
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
	sortedArr := algorithms.QuickSort([]int{1, 4, 5, 12, 3, 2, 6, 1, 4, 13}) // сортировка массива
	log.Println("========================================")
	log.Println("sorted list: ", sortedArr)
	log.Println("sorted list len: ", len(sortedArr))
	index, value, status := algorithms.BinarySearch(sortedArr, 21) // бинарный поиск
	if status == true {
		log.Printf("index: %d, value: %d ", index, value)
		log.Printf("check, get value by index: %d", sortedArr[index])
	}
	//algorithms.BFSSearchTest()
}

//func main() {
//	result := recursion.FirstRecursionFactorial(5)
//	log.Println(result)
//}
