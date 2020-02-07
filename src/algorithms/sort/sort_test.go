package sort

import (
	"fmt"
	"math"
	"sort"
	"testing"
	"time"

	"../../utils"
)

func TestBubbleSort(t *testing.T) {
	data := []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	sorted := []int{3, 5, 6, 9, 10, 17, 19, 22, 84}
	result := BubbleSort(data)
	if utils.ArrayIntEquals(sorted, result) == false {
		t.Errorf("Expected %v has %v", sorted, result)
	}
}

func TestHeapSort(t *testing.T) {
	data := []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	sorted := []int{3, 5, 6, 9, 10, 17, 19, 22, 84}
	result := HeapSort(data)
	if utils.ArrayIntEquals(sorted, result) == false {
		t.Errorf("Expected %v has %v", sorted, result)
	}
}

func TestQuicksort(t *testing.T) {
	data := []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	sorted := []int{3, 5, 6, 9, 10, 17, 19, 22, 84}
	result := Quicksort(data)
	if utils.ArrayIntEquals(sorted, result) == false {
		t.Errorf("Expected %v has %v", sorted, result)
	}
}

func StdSort(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func BenchmarkSortFuncs(b *testing.B) {
	list := []struct {
		name string
		fun  func(arr []int) []int
	}{
		{"BubbleSort", BubbleSort},
		{"HeapSort", HeapSort},
		{"Quicksort", Quicksort},
		{"StdSort", StdSort},
	}
	for _, item := range list {
		for k := 0.; k < 6; k++ {
			n := int(math.Pow(10, k))

			var data []int
			for i := n; i > 0; i-- {
				data = append(data, i)
			}

			now := time.Now()
			item.fun(data)
			done := time.Now().Sub(now)
			fmt.Printf("%s	%d	%f\n", item.name, n, done.Seconds())
		}
	}
}
