package sort

import (
	"fmt"
	"math"
	"testing"
	"time"

	"../../utils"
)

func TestBubbleSort(t *testing.T) {
	data := []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	sorted := []int{84, 22, 19, 17, 10, 9, 6, 5, 3}
	result := BubbleSort(data)
	if utils.ArrayIntEquals(sorted, result) == false {
		t.Errorf("Expected %v has %v", sorted, result)
	}
}

func TestHeapSort(t *testing.T) {
	data := []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	sorted := []int{84, 22, 19, 17, 10, 9, 6, 5, 3}
	result := HeapSort(data)
	if utils.ArrayIntEquals(sorted, result) == false {
		t.Errorf("Expected %v has %v", sorted, result)
	}
}

func BenchmarkSortFuncs(b *testing.B) {
	list := []struct {
		name string
		fun  func(arr []int) []int
	}{
		{"BubbleSort", BubbleSort},
		{"HeapSort", HeapSort},
	}
	for _, item := range list {
		for k := 0.; k < 10; k++ {
			n := int(math.Pow(4, k))

			var data []int
			for i := 1; i < n; i++ {
				data = append(data, i)
			}

			now := time.Now()
			item.fun(data)
			done := time.Now().Sub(now)
			fmt.Printf("%-10s %-10d %f seconds\n", item.name, n, done.Seconds())
		}
	}
}
