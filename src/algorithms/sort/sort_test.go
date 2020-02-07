package sort

import (
	"fmt"
	"testing"

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

func BenchmarkSearchFuncs(b *testing.B) {
	list := []struct {
		name string
		fun  func(arr []int) []int
	}{
		{"BubbleSort", BubbleSort},
		{"HeapSort", HeapSort},
	}
	for _, item := range list {
		for k := 0; k < 32000; k += 1000 {
			var data []int
			for i := 1; i < k; i++ {
				data = append(data, i)
			}
			b.Run(fmt.Sprintf("%s/%d", item.name, k), func(b *testing.B) {
				item.fun(data)
			})
		}
	}
}
