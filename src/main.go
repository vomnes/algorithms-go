package main

var graph ItemGraph

func fillGraph() *ItemGraph {
	nA := Node{HeuristicItem{"A", 6}}
	nB := Node{HeuristicItem{"B", 6}}
	nC := Node{HeuristicItem{"C", 4}}
	nD := Node{HeuristicItem{"D", 5}}
	nE := Node{HeuristicItem{"E", 4}}
	nF := Node{HeuristicItem{"F", 4}}
	nG := Node{HeuristicItem{"G", 1}}
	nH := Node{HeuristicItem{"H", 2}}
	nI := Node{"I"}
	nJ := Node{"J"}
	nK := Node{"K"}
	nL := Node{"L"}
	graph.AddNode(&nA)
	graph.AddNode(&nB)
	graph.AddNode(&nC)
	graph.AddNode(&nD)
	graph.AddNode(&nE)
	graph.AddNode(&nF)
	graph.AddNode(&nG)
	graph.AddNode(&nH)
	graph.AddNode(&nI)
	graph.AddNode(&nJ)
	graph.AddNode(&nK)
	graph.AddNode(&nL)
	// graph.AddEdge(&nA, &nB)
	// graph.AddEdge(&nA, &nC)
	// graph.AddEdge(&nB, &nE)
	// graph.AddEdge(&nC, &nE)
	// graph.AddEdge(&nC, &nF)
	// graph.AddEdge(&nE, &nF)
	// graph.AddEdge(&nD, &nA)
	// -->
	// graph.AddEdge(&nA, &nB)
	// graph.AddEdge(&nB, &nC)
	// graph.AddEdge(&nB, &nE)
	// graph.AddEdge(&nB, &nD)
	// graph.AddEdge(&nC, &nE)
	// graph.AddEdge(&nD, &nE)
	// graph.AddEdge(&nD, &nG)
	// graph.AddEdge(&nE, &nF)
	// -->
	// graph.AddEdge(&nA, &nB)
	// graph.AddEdge(&nA, &nC)
	// graph.AddEdge(&nA, &nE)
	// graph.AddEdge(&nB, &nD)
	// graph.AddEdge(&nB, &nF)
	// graph.AddEdge(&nC, &nG)
	// graph.AddEdge(&nE, &nF)
	//					A
	//				/ | \
	//		  B   C  E
	//	  /	\		|		\
	//	 D   F	G   |
	//	 		 	\		 /
	//				  --
	// graph.AddEdge(&nA, &nB)
	// graph.AddEdge(&nA, &nG)
	// graph.AddEdge(&nA, &nH)
	// graph.AddEdge(&nB, &nC)
	// graph.AddEdge(&nB, &nF)
	// graph.AddEdge(&nC, &nD)
	// graph.AddEdge(&nC, &nE)
	// graph.AddEdge(&nH, &nI)
	// graph.AddEdge(&nH, &nL)
	// graph.AddEdge(&nI, &nJ)
	// graph.AddEdge(&nI, &nK)
	//						A
	//					/ | \
	//				B		G		H
	//			/	\			/		\
	//		C		F		I			L
	//	/	\			/	\
	// D	E		J		K
	graph.AddEdge(&nA, &nB, 7)
	graph.AddEdge(&nB, &nC, 3)
	graph.AddEdge(&nC, &nD, 2)
	graph.AddEdge(&nD, &nE, 2)
	graph.AddEdge(&nD, &nF, 6)
	graph.AddEdge(&nE, &nF, 3)
	graph.AddEdge(&nE, &nG, 4)
	graph.AddEdge(&nF, &nH, 3)
	return &graph
}

func main() {
	graph := fillGraph()
	// graph.BFS()
	// fmt.Println("----")
	// graph.DFS()
	graph.AStar()
}
