package main

// type ItemGraph struct {
// 	nodes []*Node
// 	edges map[Node][]*Node
// 	lock  sync.RWMutex
// }

// BFS Travers a graph using BFS algorithm
func (g *ItemGraph) BFS() {
	g.lock.RLock()
	q := NodeQueue{}
	queue := q.New()
	queue.Enqueue(*g.nodes[0])
	visited := make(map[Node]bool)
	dp := DataPath{}
	dataPath := dp.NewListPath()
	dataPath.InsertNewPath(g.nodes[0].String())
	hasEdges := false
	for {
		if queue.IsEmpty() {
			break
		}
		currentNode := queue.Dequeue()
		visited[*currentNode] = true
		for _, childNode := range g.edges[*currentNode] {
			if !visited[*childNode] {
				queue.Enqueue(*childNode)
				visited[*childNode] = true
				dataPath.Growth(currentNode.String(), childNode.String())
				// fmt.Println(currentNode.String(), childNode.String())
				hasEdges = true
			}
		}
		if hasEdges {
			// dataPath.RemovePath(currentNode.String())
		}
	}
	dataPath.Print()
	g.lock.RUnlock()
}
