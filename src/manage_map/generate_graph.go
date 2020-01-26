package manageMap

import (
	"fmt"

	"../models"
)

func checkEdges(m *models.DataMap, prevPos, pos models.Coords) []models.Coords {
	var edges []models.Coords
	var posList = []models.Coords{
		models.Coords{Y: pos.Y - 1, X: pos.X},
		models.Coords{Y: pos.Y + 1, X: pos.X},
		models.Coords{Y: pos.Y, X: pos.X - 1},
		models.Coords{Y: pos.Y, X: pos.X + 1},
	}
	var value models.Link
	for _, checkedPos := range posList {
		value = m.GetData(checkedPos)
		// fmt.Println(value, string(value))
		if (checkedPos.Y == prevPos.Y && checkedPos.X == prevPos.X) || value.Visited {
			continue
		}
		if value.Char == models.PathKey || value.Char == models.StartKey {
			edges = append(edges, checkedPos)
		} else if value.Char == models.EndKey {
			edges = append(edges, checkedPos)
		}
	}
	// fmt.Println("#>", edges, prevPos, pos)
	return edges
}

// func lenNewEdges(m *models.DataMap, prevPos models.Coords, egdes []models.Coords) int {
// 	count := 0
// 	for _, edge range edges {
// 			if edge.Y != prevPos.Y && edge.X != prevPos.X {
// 				count++
// 			}
// 	}
// }

// GenerateGraph is the function that will grenerate our graph using map data
func GenerateGraph(dataMap *models.DataMap) {
	var currentPos, childPos, prevChildPos models.Coords
	var tmpChildPosEdges []models.Coords
	graph := models.ItemGraph{}
	g := graph.String()
	pos := dataMap.GetStart()
	// Init queue
	queue := models.NodeQueue{}
	q := queue.New()
	q.Enqueue(models.Node{Value: pos})
	// Init graph
	// countEdges := -1
	for {
		if q.IsEmpty() {
			break
		}
		currentPos = *q.Front().Value.(*models.Coords)
		// pretty.Print("@@@>", q)
		fmt.Println("&", currentPos)
		currentPosEdges := checkEdges(dataMap, models.Coords{Y: -1, X: -1}, currentPos)
		dataMap.AddEdges(currentPos, currentPosEdges)
		fmt.Println("zzzz>", currentPos, currentPosEdges)
		dataMap.PrintMap(currentPos)
		for _, edge := range currentPosEdges {
			prevChildPos = currentPos
			childPos = edge
			tmpChildPosEdges = checkEdges(dataMap, currentPos, edge)
			fmt.Println("Here:", tmpChildPosEdges, len(tmpChildPosEdges))
			for {
				dataMap.PrintMap(childPos)
				if len(tmpChildPosEdges) > 1 {
					if len(dataMap.GetEdges(childPos)) == 0 {
						dataMap.AddEdges(childPos, tmpChildPosEdges)
						q.Enqueue(models.Node{Value: &childPos})
					} else {
						fmt.Println("b")
						if !dataMap.EdgesAllVisited(childPos) {
							fmt.Println("c")
						}
					}
					break
				} else if len(tmpChildPosEdges) == 0 {
					break
				}
				dataMap.SetDataVisited(childPos, true)
				prevChildPos = childPos
				childPos = tmpChildPosEdges[0]
				tmpChildPosEdges = checkEdges(dataMap, prevChildPos, childPos)
				// fmt.Println("---->", tmpChildPosEdges)
			}
		}
		q.Dequeue()
	}
	// checkEdges(dataMap, pos.Y, pos.X)

	fmt.Print(g)
}
