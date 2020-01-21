package main

var graph ItemGraph

func fillGraph() *ItemGraph {
	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}
	nG := Node{"G"}
	graph.AddNode(&nA)
	graph.AddNode(&nB)
	graph.AddNode(&nC)
	graph.AddNode(&nD)
	graph.AddNode(&nE)
	graph.AddNode(&nF)
	graph.AddNode(&nG)

	// graph.AddEdge(&nA, &nB)
	// graph.AddEdge(&nA, &nC)
	// graph.AddEdge(&nB, &nE)
	// graph.AddEdge(&nC, &nE)
	// graph.AddEdge(&nC, &nF)
	// graph.AddEdge(&nE, &nF)
	// graph.AddEdge(&nD, &nA)
	graph.AddEdge(&nA, &nB)
	graph.AddEdge(&nB, &nA)
	graph.AddEdge(&nB, &nC)
	graph.AddEdge(&nB, &nE)
	graph.AddEdge(&nB, &nD)
	graph.AddEdge(&nC, &nE)
	graph.AddEdge(&nD, &nE)
	graph.AddEdge(&nD, &nG)
	graph.AddEdge(&nE, &nF)
	return &graph
}

func main() {
	graph := fillGraph()
	graph.BFS()
}
