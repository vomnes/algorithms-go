package manageMap

// Exec exec
func Exec() {
	m := storeInputMap("/Users/vomnes/Documents/programming/graph-algorithms-go/input/maze.map")
	// m := storeInputMap("/Users/vomnes/Documents/programming/graph-algorithms-go/input/maze1.map")
	// m := storeInputMap("/Users/vomnes/Documents/programming/graph-algorithms-go/input/maze2.map")
	// m := storeInputMap("/Users/vomnes/Documents/programming/graph-algorithms-go/input/maze3.map")
	graph := GenerateGraph(m)
	graph.Print()
	// algorithms.BFS(graph)
	// pretty.Print(m)
}
