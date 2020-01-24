package algorithms

import (
	"fmt"

	"../models"
)

// ItemGraph the Items graph
// type ItemGraph struct {
// 	nodes []*Node
// 	edges map[Node][]*Node
// 	lock  sync.RWMutex
// }

// DFS is the DFS algorithm implementation
func DFS(g *models.ItemGraph) {
	var s *models.NodeStack
	var currentNode *models.Node
	// hasEdges is a boolean that allows to check if we need to remove the
	// previous paths that have been updated
	var hasEdges bool
	// Init stack
	stack := models.NodeStack{}
	s = stack.New()
	s.Push(*g.GetStart())
	// Visited map
	visited := make(map[models.Node]bool)
	// Init Path List
	dp := models.DataPath{}
	dataPath := dp.NewListPath()
	dataPath.InsertNewPath(g.GetStart().String())
	for {
		if stack.IsEmpty() {
			break
		}
		currentNode = s.Pop()
		hasEdges = false
		if !visited[*currentNode] {
			visited[*currentNode] = true
			fmt.Println("Node visited:", currentNode)
			for _, childNode := range g.GetEdges(currentNode) {
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
}
