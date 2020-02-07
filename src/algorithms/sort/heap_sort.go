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
	// Create the heap
	heap := models.Heap{}
	h := heap.New(len, models.HeapMin)
	// Add nodes with the arr values in the heap
	for _, e := range arr {
		h.Insert(e)
	}
	// Allocat sorted array
	sortedArray := make([]int, h.GetCapacity())
	index := 0
	for {
		// Store the root value then remove it
		sortedArray[index] = h.Delete(0)
		if h.IsEmpty() {
			break
		}
		index++
	}
	h = nil
	return sortedArray
}
