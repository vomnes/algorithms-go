package sort

import (
	"../../models"
)

// HeapSort allows to sort an int array using heap data structure
func HeapSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	len := len(arr)
	heap := models.Heap{}
	h := heap.New(len, models.HeapMax)
	for _, e := range arr {
		h.Insert(e)
	}
	sortedArray := make([]int, h.GetCapacity())
	index := 0
	for {
		sortedArray[index] = h.DeleteRoot()
		if h.IsEmpty() {
			break
		}
		index++
	}
	return sortedArray
}
