package models

import (
	"fmt"

	"../utils"
)

// Heap is the structure of the heap data structure
type Heap struct {
	array           []int
	count, capacity int
	kind            heapType
}

type heapType int

const (
	// HeapMin is used to set the heap min type
	HeapMin heapType = 0
	// HeapMax is used to set the heap max type
	HeapMax heapType = 1
)

func (h *Heap) new(capacity int, kind heapType) *Heap {
	h.array = make([]int, capacity)
	h.capacity = capacity
	h.count = 0
	h.kind = kind
	return h
}

func (h *Heap) insert(value int) {
	if h.count >= h.capacity {
		return
	}
	h.array[h.count] = value
	if h.kind == HeapMin {
		h.heapifyMin(h.count)
	} else if h.kind == HeapMax {
		// WORK IN PROGRESS
		for i := 0; i < h.count; i++ {
			fmt.Print(i)
			h.heapifyMax(i)
		}
		fmt.Print("\n")
	}
	h.count++
}

func (h *Heap) heapifyMin(nodeIndex int) {
	parentNodeIndex := (nodeIndex - 1) / 2
	if h.array[nodeIndex] < h.array[parentNodeIndex] {
		utils.SwapInt(&h.array[nodeIndex], &h.array[parentNodeIndex])
		h.heapifyMin(parentNodeIndex)
	}
}

func (h *Heap) heapifyMax(nodeIndex int) {
	left := nodeIndex*2 + 1
	right := nodeIndex*2 + 2
	largest := nodeIndex

	if left <= h.count && h.array[left] > h.array[largest] {
		largest = left
	}
	if right <= h.count && h.array[right] > h.array[largest] {
		largest = right
	}
	if largest != nodeIndex {
		utils.SwapInt(&h.array[nodeIndex], &h.array[largest])
		h.heapifyMax(largest)
	}
}

// ExecHeap exec
func ExecHeap() {
	data := []int{19, 2, 17, 36, 1, 7, 3, 100, 25}
	len := len(data)
	heap := Heap{}
	h := heap.new(len, HeapMax)
	for _, e := range data {
		h.insert(e)
	}
	fmt.Println(h)
}
