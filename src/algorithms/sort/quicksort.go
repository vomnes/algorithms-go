package sort

import (
	"../../utils"
)

type quicksort struct {
	array   []int
	cmpFunc func(a, b int) bool
}

func cmp(a, b int) bool {
	return a < b
}

func (d *quicksort) partition(lo, hi int) int {
	pivot := d.array[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if d.cmpFunc(d.array[j], pivot) {
			utils.SwapInt(&d.array[i], &d.array[j])
			i++
		}
	}
	utils.SwapInt(&d.array[i], &d.array[hi])
	return i
}

func (d *quicksort) qs(lo, hi int) {
	if lo < hi {
		p := d.partition(lo, hi)
		d.qs(lo, p-1)
		d.qs(p+1, hi)
	}
}

// Quicksort is the quick sort algorithm implementation
// Time complexity worst case : O(n^2)
func Quicksort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	data := quicksort{}
	data.array = arr
	data.cmpFunc = cmp
	data.qs(0, len(arr)-1)
	return data.array
}
