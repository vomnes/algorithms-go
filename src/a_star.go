package main

import "fmt"

// AStar is the A* algorithm implementation
func (g *ItemGraph) AStar() {
	var currentNode *Node
	var tentativeGScore int

	aStarData := AStarData{}
	d := aStarData.New()
	d.AddItem(*g.nodes[0])
	d.SetGScore(*g.nodes[0], 0)
	d.SetFScore(*g.nodes[0], g.nodes[0].value.(HeuristicItem).heuristic)
	for {
		if d.OpenSetIsEmpty() {
			break
		}
		currentNode = d.GetLowestFScoreNodeInOpenSet()
		// if current = goal
		// 	return reconstruct_path(cameFrom, current)
		fmt.Println("->", currentNode)
		d.RemoveItemFromOpenSet(*currentNode)
		for indexChildNode, childNode := range g.edges[*currentNode] {
			fmt.Println("Child Nodes:", childNode)
			tentativeGScore = d.GetGScore(*currentNode) + g.GetDistance(currentNode, indexChildNode)
			fmt.Println(d.GetGScore(*currentNode), g.GetDistance(currentNode, indexChildNode), "-", d.GetGScore(*childNode))
			if tentativeGScore < d.GetGScore(*childNode) || d.items[*childNode] == (elem{}) {
				d.SetComeFrom(*childNode, *currentNode)
				d.SetGScore(*childNode, tentativeGScore)
				d.SetFScore(*childNode, d.GetGScore(*childNode)+g.GetHeuristicValue(childNode))
				d.SetInOpenSet(*childNode)
				fmt.Println(childNode.value.(HeuristicItem).data)
			}
		}
	}
	// g.Print()
}
