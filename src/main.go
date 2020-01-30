package main

import (
	"fmt"

	"./algorithms"
	"./models"
)

var graph models.ItemGraph

func fillGraph() *models.ItemGraph {
	// nA := Node{HeuristicItem{"A", 6}}
	// nB := Node{HeuristicItem{"B", 6}}
	// nC := Node{HeuristicItem{"C", 4}}
	// nD := Node{HeuristicItem{"D", 5}}
	// nE := Node{HeuristicItem{"E", 4}}
	// nF := Node{HeuristicItem{"F", 4}}
	// nG := Node{HeuristicItem{"G", 1}}
	// nH := Node{HeuristicItem{"H", 2}}
	// nA := models.Node{models.HeuristicItem{"A", 10}}
	// nB := models.Node{models.HeuristicItem{"B", 8}}
	// nC := models.Node{models.HeuristicItem{"C", 5}}
	// nD := models.Node{models.HeuristicItem{"D", 7}}
	// nE := models.Node{models.HeuristicItem{"E", 3}}
	// nF := models.Node{models.HeuristicItem{"F", 6}}
	// nG := models.Node{models.HeuristicItem{"G", 5}}
	// nH := models.Node{models.HeuristicItem{"H", 3}}
	// nI := models.Node{models.HeuristicItem{"I", 1}}
	// nJ := models.Node{models.HeuristicItem{"J", 0}}
	// nK := models.Node{"K"}
	// nL := models.Node{"L"}
	// graph.AddNode(&nA)
	// graph.AddNode(&nB)
	// graph.AddNode(&nC)
	// graph.AddNode(&nD)
	// graph.AddNode(&nE)
	// graph.AddNode(&nF)
	// graph.AddNode(&nG)
	// graph.AddNode(&nH)
	// graph.AddNode(&nI)
	// graph.AddNode(&nJ)
	// graph.AddNode(&nK)
	// graph.AddNode(&nL)
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
	// -->
	// graph.SetStart(&nA)
	// graph.AddEdge(&nA, &nB, 7)
	// graph.AddEdge(&nB, &nC, 3)
	// graph.AddEdge(&nC, &nD, 2)
	// graph.AddEdge(&nD, &nE, 2)
	// graph.AddEdge(&nD, &nF, 6)
	// graph.AddEdge(&nE, &nF, 3)
	// graph.AddEdge(&nE, &nG, 4)
	// graph.AddEdge(&nF, &nH, 3)
	// graph.SetEnd(&nG)
	// -->
	// graph.SetStart(&nA)
	// graph.AddEdge(&nA, &nB, 6)
	// graph.AddEdge(&nA, &nF, 3)
	// graph.AddEdge(&nB, &nC, 3)
	// graph.AddEdge(&nB, &nD, 2)
	// graph.AddEdge(&nC, &nD, 1)
	// graph.AddEdge(&nC, &nE, 5)
	// graph.AddEdge(&nD, &nE, 8)
	// graph.AddEdge(&nE, &nJ, 5)
	// graph.AddEdge(&nF, &nG, 1)
	// graph.AddEdge(&nF, &nH, 7)
	// graph.AddEdge(&nG, &nI, 3)
	// graph.AddEdge(&nH, &nI, 2)
	// graph.AddEdge(&nI, &nJ, 3)
	// graph.SetEnd(&nJ)
	return &graph
}

func main() {
	// graph := fillGraph()
	// algorithms.BFS(graph)
	// fmt.Println("----")
	// algorithms.DFS(graph)
	// algorithms.AStar(graph)
	// _map.Exec()

	fmt.Println(
		algorithms.MaxXor(
			[]int32{1, 3, 5, 7},
			[]int32{17, 6},
		),
	)
}
