package main

import "fmt"

func (d *AStarData) reconstructPath(startNode, lastNode *Node) []string {
	var currentNode Node
	path := []string{}
	currentNode = *lastNode
	for {
		path = append([]string{currentNode.value.(HeuristicItem).data}, path...)
		if currentNode.value == startNode.value {
			break
		}
		currentNode = d.items[currentNode].cameFrom
	}
	return path
}

// AStar is the A* algorithm implementation
func (g *ItemGraph) AStar() {
	var currentNode *Node
	var tentativeGScore int
	// Init A Star Data
	aStarData := AStarData{}
	d := aStarData.New()
	d.Init(g.nodes)
	start := *g.nodes[0]
	d.SetInOpenSet(start)
	d.SetGScore(start, 0)
	d.SetFScore(start, start.value.(HeuristicItem).heuristic)
	for {
		if d.OpenSetIsEmpty() {
			break
		}
		currentNode = d.GetLowestFScoreNodeInOpenSet()
		if g.IsEndNode(currentNode) {
			fmt.Println(d.reconstructPath(g.GetStart(), g.GetEnd()))
			break
		}
		d.RemoveItemFromOpenSet(*currentNode)
		for indexChildNode, childNode := range g.edges[*currentNode] {
			tentativeGScore = d.GetGScore(*currentNode) + g.GetDistance(currentNode, indexChildNode)
			if tentativeGScore < d.GetGScore(*childNode) {
				d.SetComeFrom(*childNode, *currentNode)
				d.SetGScore(*childNode, tentativeGScore)
				d.SetFScore(*childNode, d.GetGScore(*childNode)+g.GetHeuristicValue(childNode))
				d.SetInOpenSet(*childNode)
			}
		}
	}
}
