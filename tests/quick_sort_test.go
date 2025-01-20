package tests

import (
	"learn/algorithms"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testList := []int{1, 4, 5, 12, 3, 2, 6, 1, 4, 13}
	expectedList := []int{1, 1, 2, 3, 4, 4, 5, 6, 12, 13}
	result := algorithms.QuickSort(testList)
	for index, value := range result {
		if value != expectedList[index] {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedList)
		}
	}
}
