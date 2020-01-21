package main

var graph ItemGraph

func fillGraph() *ItemGraph {
	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}
	graph.AddNode(&nA)
	graph.AddNode(&nB)
	graph.AddNode(&nC)
	graph.AddNode(&nD)
	graph.AddNode(&nE)
	graph.AddNode(&nF)

	graph.AddEdge(&nA, &nB)
	graph.AddEdge(&nA, &nC)
	graph.AddEdge(&nB, &nE)
	graph.AddEdge(&nC, &nE)
	graph.AddEdge(&nC, &nF)
	graph.AddEdge(&nE, &nF)
	graph.AddEdge(&nD, &nA)
	return &graph
}

func main() {
	graph := fillGraph()
	graph.Print()
	// graph.BFS()

	dp := DataPath{}
	dataPath := dp.NewListPath()
	firstList := dataPath.InsertNewPath("C")
	dataPath.CopyPathAddNewItem(*firstList, "B")
	dataPath.Print()
	// pretty.Print(dataPath.listPath.Back())
}
