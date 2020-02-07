package sort

import (
	"../../utils"
)

// BubbleSort is the implementation of the bubble sort algorithm
func BubbleSort(arr []int) []int {
	len := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for index := 0; index < len-1; index++ {
			if arr[index] > arr[index+1] {
				utils.SwapInt(&arr[index], &arr[index+1])
				swapped = true
			}
		}
	}
	return arr
}
