package algorithms

// модуль с рекурсивной функцией быстрой сортировки

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[(len(list)-1)/2] // опорный элемент
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
	result := append(QuickSort(less), pivot)       // создаем массив, сортируем меньший подмассив + екстендим pivot'ом
	result = append(result, QuickSort(greater)...) // добавлям к нему отсортированный больший подмассив

	return result
}
