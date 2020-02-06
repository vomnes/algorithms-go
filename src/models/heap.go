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
		h.heapify(h.count, minOrder)
	} else if h.kind == HeapMax {
		h.heapify(h.count, maxOrder)
	}
	h.count++
}

func minOrder(a, b int) bool {
	return a < b
}

func maxOrder(a, b int) bool {
	return a > b
}

func (h *Heap) heapify(nodeIndex int, cmp func(int, int) bool) {
	parentNodeIndex := (nodeIndex - 1) / 2
	if cmp(h.array[nodeIndex], h.array[parentNodeIndex]) {
		utils.SwapInt(&h.array[nodeIndex], &h.array[parentNodeIndex])
		h.heapify(parentNodeIndex, cmp)
	}
}

// func (h *Heap) down(nodeIndex int) {
// 	leftChild := nodeIndex * 2
// 	rightChild := nodeIndex*2 + 1
// 	// Stop if no childs (if no left child then no right child)
// 	if leftChild > h.count || rightChild > h.count {
// 		return
// 	}
//
// 	if h.array[leftChild] > h.array[rightChild] {
// 		utils.SwapInt(&h.array[leftChild], &h.array[nodeIndex])
// 		h.down(leftChild)
// 	} else {
// 		utils.SwapInt(&h.array[rightChild], &h.array[nodeIndex])
// 		h.down(rightChild)
// 	}
// }

// ExecHeap exec
func ExecHeap() {
	data := []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	len := len(data)
	heap := Heap{}
	h := heap.new(len, HeapMin)
	fmt.Println(h)
	for _, e := range data {
		h.insert(e)
	}
	fmt.Println(h)

}
