package algorithms

import (
	"fmt"
	"math"
	"testing"
)

var limit = 10000000
var data []int

func BenchmarkMain(b *testing.B) {
	for i := 0; i < limit; i++ {
		data = append(data, i)
	}
}

func BenchmarkSearchFuncs(b *testing.B) {
	list := []struct {
		name string
		fun  func(data []int, length, toFind int) bool
	}{
		{"BinarySearch", BinarySearch},
		{"IterativeSearch", IterativeSearch},
	}
	for _, item := range list {
		for k := 0.; k <= 10; k++ {
			n := int(math.Pow(2, k))
			b.Run(fmt.Sprintf("%s/%d", item.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					item.fun(data, len(data), n)
				}
			})
		}
	}
}

func TestIterativeSearch(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5}
	result := IterativeSearch(data, len(data), 2)
	if result == false {
		t.Error("Expected true has false")
	}
	result = IterativeSearch(data, len(data), 100)
	if result == true {
		t.Error("Expected false has true")
	}
}

func TestBinarySearch(t *testing.T) {
	data := []int{0, 1, 2, 3, 4, 5}
	result := BinarySearch(data, len(data), 2)
	if result == false {
		t.Error("Expected true has false")
	}
	result = BinarySearch(data, len(data), 3)
	if result == false {
		t.Error("Expected true has false")
	}
	result = BinarySearch(data, len(data), 100)
	if result == true {
		t.Error("Expected false has true")
	}
}
