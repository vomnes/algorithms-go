package algorithms

import (
	"fmt"

	"../models"
)

// type ItemGraph struct {
// 	nodes []*Node
// 	edges map[Node][]*Node
// 	lock  sync.RWMutex
// }

// BFS is the BFS algorithm implementation
func BFS(g *models.ItemGraph) {
	// hasEdges is a boolean that allows to check if we need to remove the
	// previous paths that have been updated
	var hasEdges bool
	visited := make(map[models.Node]bool)
	// Init queue
	q := models.NodeQueue{}
	queue := q.New()
	queue.Enqueue(*g.GetStart())
	// Init Path List
	dp := models.DataPath{}
	dataPath := dp.NewListPath()
	dataPath.InsertNewPath(g.GetStart().String())
	for {
		if queue.IsEmpty() {
			break
		}
		currentNode := queue.Dequeue()
		visited[*currentNode] = true
		hasEdges = false
		fmt.Println("Node visited:", currentNode)
		for _, childNode := range g.GetEdges(currentNode) {
			if !visited[*childNode] {
				queue.Enqueue(*childNode)
				visited[*childNode] = true
				dataPath.Growth(currentNode.String(), childNode.String())
				hasEdges = true
			}
		}
		if hasEdges {
			dataPath.RemovePath(currentNode.String())
		}
	}
	dataPath.Print()
}
