package main

import (
	"fmt"
	"log"
	"math"
	"sync"
)

type score struct {
	g int
	f int
}

type elem struct {
	inOpenSet bool
	score     score
	cameFrom  Node
}

// AStarData store the  A Star data
type AStarData struct {
	items map[Node]elem
	lock  sync.RWMutex
}

// New creates a new A Star data
func (s *AStarData) New() *AStarData {
	s.lock.Lock()
	s.items = make(map[Node]elem)
	s.lock.Unlock()
	return s
}

// AddItem adds a new elem to the A Star data
func (s *AStarData) AddItem(e Node) {
	s.lock.Lock()
	s.items[e] = elem{
		inOpenSet: true,
		score: score{
			g: math.MaxInt32,
			f: math.MaxInt32,
		},
	}
	s.lock.Unlock()
}

// RemoveItemFromOpenSet set the inOpenSet to false of a given Node
func (s *AStarData) RemoveItemFromOpenSet(e Node) {
	s.lock.Lock()
	data := s.items[e]
	data.inOpenSet = false
	s.items[e] = data
	s.lock.Unlock()
}

// SetFScore set the fScore value of the selected node
func (s *AStarData) SetFScore(e Node, value int) {
	s.lock.Lock()
	data := s.items[e]
	data.score.f = value
	s.items[e] = data
	s.lock.Unlock()
}

// SetGScore set the gScore value of the selected node
func (s *AStarData) SetGScore(e Node, value int) {
	s.lock.Lock()
	data := s.items[e]
	data.score.g = value
	s.items[e] = data
	s.lock.Unlock()
}

// GetGScore return the gScore value of the selected node
func (s *AStarData) GetGScore(e Node) int {
	return s.items[e].score.g
}

// SetComeFrom set the comeFrom value of the selected node
func (s *AStarData) SetComeFrom(e Node, value Node) {
	s.lock.Lock()
	data := s.items[e]
	data.cameFrom = value
	s.items[e] = data
	s.lock.Unlock()
}

// SetInOpenSet set the inOpenSet value of the selected node
func (s *AStarData) SetInOpenSet(e Node) {
	s.lock.Lock()
	data := s.items[e]
	data.inOpenSet = true
	s.items[e] = data
	s.lock.Unlock()
}

// GetLowestFScoreNodeInOpenSet ...
func (s *AStarData) GetLowestFScoreNodeInOpenSet() *Node {
	var node *Node
	tmpMin := math.MaxInt32
	for checkedNode, elem := range s.items {
		if elem.inOpenSet == true {
			if elem.score.f <= tmpMin {
				tmpMin = elem.score.f
				node = &checkedNode
			}
		}
	}
	if node == nil {
		log.Fatal("GetLowestFScoreNodeInOpenSet: Node is nil")
	}
	return node
}

// Print print the a star data
func (s *AStarData) Print() {
	fmt.Println("AStarData:", s.items)
}

// OpenSetIsEmpty return true if all the inOpenSet values are false
func (s *AStarData) OpenSetIsEmpty() bool {
	for _, elem := range s.items {
		if elem.inOpenSet == true {
			return false
		}
	}
	return true
}
