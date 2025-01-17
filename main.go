package main

import (
	"log"
	"math/rand"
)

func binarySearch(list []int, item int) (int, bool) {
	low := 0              // левая граница
	high := len(list) - 1 // правая граница
	for low <= high {
		mid := (low + high) / 2 // середина списка
		guess := list[mid]      // элемент в середине списка
		switch {
		case guess == item:
			return mid, true
		case guess > item:
			high = mid - 1 // сдвигаем правую границу на середину
		default:
			low = mid + 1 // сдвигаем левую границу на середину
		}
	}
	return 0, false
}

func createArray() []int {
	myArray := make([]int, 100)
	//rand.Intn(200)
	for i := range myArray {
		myArray[i] = rand.Intn(211)
	}
	log.Println("list len", len(myArray))
	return myArray
}

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[(len(list)-1)/1] // опорный индекс
	var less []int
	var greater []int
	for _, value := range list {
		if value < pivot {
			less = append(less, value) // выделяем подмассив меньше pivot
		}
	}
	for _, value := range list {
		if value > pivot {
			greater = append(greater, value) // выделяем подмассив больше pivot
		}
	}
	result := append(quickSort(less), pivot) // создаем массив, сортируем меньший подмассив + екстендим pivot'ом
	result = append(result, quickSort(greater)...) // добавлям к нему отсортированный больший подмассив

	return result
}

func main() {
	arr := createArray()
	quickSort(arr)
	//log.Println("done", sorted)
}
