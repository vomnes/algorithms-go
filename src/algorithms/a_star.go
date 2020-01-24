package algorithms

import "fmt"
import models "../models"

func reconstructPath(d *models.AStarData, startNode, lastNode *models.Node) []string {
	var currentNode models.Node
	path := []string{}
	currentNode = *lastNode
	for {
		path = append([]string{currentNode.Value.(models.HeuristicItem).Data}, path...)
		if currentNode.Value == startNode.Value {
			break
		}
		currentNode = d.GetCameFrom(&currentNode)
	}
	return path
}

// AStar is the A* algorithm implementation
func AStar(g *models.ItemGraph) {
	var currentNode *models.Node
	var tentativeGScore int
	// Init A Star Data
	aStarData := models.AStarData{}
	d := aStarData.New()
	d.Init(g.GetNodes())
	start := *g.GetStart()
	d.SetInOpenSet(start)
	d.SetGScore(start, 0)
	d.SetFScore(start, start.Value.(models.HeuristicItem).Heuristic)
	for {
		if d.OpenSetIsEmpty() {
			break
		}
		currentNode = d.GetLowestFScoreNodeInOpenSet()
		if g.IsEndNode(currentNode) {
			fmt.Println(reconstructPath(d, g.GetStart(), g.GetEnd()))
			break
		}
		d.RemoveItemFromOpenSet(*currentNode)
		for indexChildNode, childNode := range g.GetEdges(currentNode) {
			tentativeGScore = d.GetGScore(*currentNode) + g.GetDistance(currentNode, indexChildNode)
			if tentativeGScore < d.GetGScore(*childNode) {
				d.SetCameFrom(*childNode, *currentNode)
				d.SetGScore(*childNode, tentativeGScore)
				d.SetFScore(*childNode, d.GetGScore(*childNode)+g.GetHeuristicValue(childNode))
				d.SetInOpenSet(*childNode)
			}
		}
	}
}
