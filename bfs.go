package main

import "fmt"

// type ItemGraph struct {
// 	nodes []*Node
// 	edges map[Node][]*Node
// 	lock  sync.RWMutex
// }

// BFS Travers a graph using BFS algorithm
func (g *ItemGraph) BFS() {
	g.lock.RLock()
	paths := make([][]string, 0)

	q := NodeQueue{}
	queue := q.New()
	queue.Enqueue(*g.nodes[0])
	visited := make(map[Node]bool)
	for {
		if queue.IsEmpty() {
			break
		}
		currentNode := queue.Dequeue()
		visited[*currentNode] = true
		paths = append(paths, []string{currentNode.String()})
		for _, childNode := range g.edges[*currentNode] {
			fmt.Println(childNode)
			if !visited[*childNode] {
				queue.Enqueue(*childNode)
				visited[*childNode] = true
			}
		}
	}
	// fmt.Println(paths)
	g.lock.RUnlock()
}
