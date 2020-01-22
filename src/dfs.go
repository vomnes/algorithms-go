package main

import (
	"fmt"
)

// ItemGraph the Items graph
// type ItemGraph struct {
// 	nodes []*Node
// 	edges map[Node][]*Node
// 	lock  sync.RWMutex
// }

// DFS is the DFS algorithm implementation
func (g *ItemGraph) DFS() {
	var s *NodeStack
	var currentNode *Node
	g.lock.RLock()
	// hasEdges is a boolean that allows to check if we need to remove the
	// previous paths that have been updated
	var hasEdges bool
	// Init stack
	stack := NodeStack{}
	s = stack.New()
	s.Push(*g.nodes[0])
	// Visited map
	visited := make(map[Node]bool)
	// Init Path List
	dp := DataPath{}
	dataPath := dp.NewListPath()
	dataPath.InsertNewPath(g.nodes[0].String())
	for {
		if stack.IsEmpty() {
			break
		}
		currentNode = s.Pop()
		hasEdges = false
		if !visited[*currentNode] {
			visited[*currentNode] = true
			fmt.Println("Node visited:", currentNode)
			for _, childNode := range g.edges[*currentNode] {
				s.Push(*childNode)
				if !visited[*childNode] {
					dataPath.Growth(currentNode.String(), childNode.String())
					hasEdges = true
				}
			}
		}
		if hasEdges {
			dataPath.RemovePath(currentNode.String())
		}
	}
	dataPath.Print()
	g.lock.RUnlock()
}
