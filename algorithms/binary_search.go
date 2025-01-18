package algorithms

// модуль с функцией бинарного поиска

func BinarySearch(list []int, item int) (int, int, bool) {
	low := 0              // левая граница
	high := len(list) - 1 // правая граница
	for low <= high {
		mid := (low + high) / 2 // середина списка
		guess := list[mid]      // элемент в середине списка
		switch {
		case guess == item:
			return mid, guess, true
		case guess > item:
			high = mid - 1 // сдвигаем правую границу на середину
		default:
			low = mid + 1 // сдвигаем левую границу на середину
		}
	}
	return 0, 0, false
}
