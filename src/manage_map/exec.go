package manageMap

// Exec exec
func Exec() {
	m := storeInputMap("/Users/vomnes/Documents/programming/graph-algorithms-go/input/maze.map")
	GenerateGraph(m)
	// pretty.Print(m)
}
