package main

import "sync"

// NodeStack is the stack of Nodes
type NodeStack struct {
	items []Node
	lock  sync.RWMutex
}

// New creates a new NodeStack
func (s *NodeStack) New() *NodeStack {
	s.lock.Lock()
	s.items = []Node{}
	s.lock.Unlock()
	return s
}

// Push adds an Node to the end of the stack
func (s *NodeStack) Push(t Node) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Pop removes an Node from the end of the stack
func (s *NodeStack) Pop() *Node {
	s.lock.Lock()
	var item Node
	item = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.lock.Unlock()
	return &item
}

// IsEmpty returns true if the stack is empty
func (s *NodeStack) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items) == 0
}

// Size returns the number of Nodes in the stack
func (s *NodeStack) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.items)
}
