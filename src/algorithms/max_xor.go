package algorithms

import "fmt"

// BinaryNode is a node of the binary tree
type BinaryNode struct {
	// left  -> 0
	// right -> 1
	left, right *BinaryNode
	data        interface{}
}

// BinaryTree is the structure of the binary tree
type BinaryTree struct {
	root *BinaryNode
}

func getBit(nb int32, index int) int {
	return int((nb >> index) & 1)
}

func setBit(nb int32, index int) int32 {
	return (1 << index) | nb
}

func (s *BinaryTree) insert(nb int32) {
	var bit int
	currentNode := s.root
	for i := 31; i >= 0; i-- {
		bit = getBit(nb, i)
		if bit == 0 {
			// Bit 0
			if currentNode.left == nil {
				currentNode.left = &BinaryNode{left: nil, right: nil, data: bit}
			}
			currentNode = currentNode.left
		} else {
			// Bit 1
			if currentNode.right == nil {
				currentNode.right = &BinaryNode{left: nil, right: nil, data: bit}
			}
			currentNode = currentNode.right
		}
	}
}

func (s *BinaryTree) new() *BinaryTree {
	s.root = &BinaryNode{left: nil, right: nil, data: nil}
	return s
}

func (s *BinaryTree) print(n *BinaryNode, space int) {
	if n == nil {
		return
	}
	if n.right != nil {
		s.print(n.right, space+1)
	}
	fmt.Print("\t")
	for i := 0; i <= space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(n.data)
	if n.left != nil {
		s.print(n.left, space+1)
	}
}

func (s *BinaryTree) getMaxXor(toXor int32) int32 {
	var xorOutput int
	var max int32
	currentNode := s.root
	// For each bit of toXor
	for index := 31; index >= 0; index-- {
		xorOutput = getBit(toXor, index)
		// Check which node 0 or 1 return the higher XOR result
		if xorOutput^0 == 1 {
			if currentNode.left != nil {
				currentNode = currentNode.left
				max = setBit(max, index)
			} else {
				// Zero node at this level doesn't exists so next one is right
				// We don't need to set the bit because he is equal to zero
				currentNode = currentNode.right
			}
		} else if xorOutput^1 == 1 {
			if currentNode.right != nil {
				currentNode = currentNode.right
				max = setBit(max, index)
			} else {
				// One node at this level doesn't exists so next one is left
				// We don't need to set the bit because he is equal to zero
				currentNode = currentNode.left
			}
		}
	}
	return max
}

// MaxXor will determines the max xor of a given arr for each element of the
// queries array
func MaxXor(arr, queries []int32) []int32 {
	var result []int32
	tree := BinaryTree{}
	t := tree.new()
	for _, elem := range arr {
		t.insert(elem)
	}
	for _, elem := range queries {
		result = append(result, t.getMaxXor(elem))
	}
	return result
}
