package models

import (
	"fmt"

	"../utils"
)

// Heap is the structure of the heap data structure
type Heap struct {
	array           []int
	count, capacity int
	cmpFunc         func(int, int) bool
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
	if kind == HeapMin {
		h.cmpFunc = minOrder
	} else if kind == HeapMax {
		h.cmpFunc = maxOrder
	}
	return h
}

func (h *Heap) insert(value int) {
	if h.count >= h.capacity {
		return
	}
	h.array[h.count] = value
	h.heapify(h.count, h.cmpFunc)
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

func (h *Heap) pop() {
	h.capacity--
	h.array = h.array[:h.capacity]
	h.count--
}

func (h *Heap) delete(nodeIndex int) {
	lastNodeIndex := h.count - 1
	if lastNodeIndex != nodeIndex {
		utils.SwapInt(&h.array[lastNodeIndex], &h.array[nodeIndex])
		h.down(0)
	}
	h.pop()
}

func (h *Heap) down(nodeIndex int) {
	leftChild := nodeIndex * 2
	rightChild := nodeIndex*2 + 1
	extremity := nodeIndex
	// Stop if no childs (if no left child then no right child)
	if leftChild > h.count || leftChild < 0 {
		return
	}
	if h.cmpFunc(h.array[leftChild], h.array[nodeIndex]) {
		extremity = leftChild
	}
	if rightChild <= h.count && rightChild >= 0 &&
		h.cmpFunc(h.array[rightChild], h.array[nodeIndex]) {
		extremity = rightChild
	}
	if extremity != nodeIndex {
		utils.SwapInt(&h.array[extremity], &h.array[nodeIndex])
		h.down(extremity)
	}
	return
}

// ExecHeap exec
func ExecHeap() {
	data := []int{5, 3, 17, 10, 84, 19, 6, 22, 9}
	len := len(data)
	heap := Heap{}
	h := heap.new(len, HeapMax)
	for _, e := range data {
		h.insert(e)
	}
	fmt.Println(h)
	h.delete(0)
	h.delete(0)
	h.delete(0)
	fmt.Println(h)
}
