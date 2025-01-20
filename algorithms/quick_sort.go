package algorithms

// модуль с рекурсивной функцией быстрой сортировки

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[(len(list)-1)/2] // опорный элемент, лучше брать случайный, тк первый элемент плохо работает
	// при отсортированном массиве, медиана вычисляется трудозатрат
	var less []int
	var greater []int
	var equal []int
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
	for _, value := range list {
		if value == pivot {
			equal = append(equal, value) // выделяем подмассив равному pivot(что бы сохранить дубли)
		}
	}
	sortedLess := QuickSort(less)
	list = append(sortedLess, equal...) // создаем массив, сортируем меньший подмассив + екстендим pivot'ом
	sortedGreater := QuickSort(greater)
	list = append(list, sortedGreater...) // добавляем к нему отсортированный больший подмассив

	return list
}

func QuickSortTest(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[(len(list)-1)/2] // опорный элемент, лучше брать случайный, тк первый элемент плохо работает
	// при отсортированном массиве, медиана вычисляется трудозатрат
	var less []int
	var greater []int
	var equal []int
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
	for _, value := range list {
		if value == pivot {
			equal = append(equal, value) // выделяем подмассив равному pivot(что бы сохранить дубли)
		}
	}
	sortedLess := QuickSortTest(less)
	list = append(sortedLess, equal...) // создаем массив, сортируем меньший подмассив + екстендим pivot'ом
	return list
}

func QuickSortGreater(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[(len(list)-1)/2] // опорный элемент, лучше брать случайный, тк первый элемент плохо работает
	// при отсортированном массиве, медиана вычисляется трудозатрат
	var less []int
	var greater []int
	var equal []int
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
	for _, value := range list {
		if value == pivot {
			equal = append(equal, value) // выделяем подмассив равному pivot(что бы сохранить дубли)
		}
	}
	sortedGreater := QuickSortGreater(greater)
	list = append(sortedGreater, equal...) // добавляем к нему отсортированный больший подмассив
	return list
}
