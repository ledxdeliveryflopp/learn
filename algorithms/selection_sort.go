package algorithms

func deleteElemByIndex(list []int, index int) []int {
	list = append(list[:index], list[index+1:]...)
	return list
}

func findSmallest(list []int) int {
	smallest := list[0]
	smallestIndex := 0
	for index, value := range list {
		if value < smallest {
			smallest = value
			smallestIndex = index
		}
	}
	return smallestIndex
}

func SelectionSort(list []int) []int {
	var newList []int
	for _, _ = range list {
		smallest := findSmallest(list)
		newList = append(newList, list[smallest])
		list = deleteElemByIndex(list, smallest)
	}
	return newList
}
