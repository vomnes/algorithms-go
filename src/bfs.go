package main

// type ItemGraph struct {
// 	nodes []*Node
// 	edges map[Node][]*Node
// 	lock  sync.RWMutex
// }

// BFS Travers a graph using BFS algorithm
func (g *ItemGraph) BFS() {
	// hasEdges is a boolean that allows to check if we need to remove the
	// previous paths that have been updated
	var hasEdges bool
	g.lock.RLock()
	visited := make(map[Node]bool)
	// Init queue
	q := NodeQueue{}
	queue := q.New()
	queue.Enqueue(*g.nodes[0])
	// Init Path List
	dp := DataPath{}
	dataPath := dp.NewListPath()
	dataPath.InsertNewPath(g.nodes[0].String())
	for {
		if queue.IsEmpty() {
			break
		}
		currentNode := queue.Dequeue()
		visited[*currentNode] = true
		hasEdges = false
		for _, childNode := range g.edges[*currentNode] {
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
	g.lock.RUnlock()
}
