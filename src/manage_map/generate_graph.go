package manageMap

import (
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
		if checkedPos.Y == prevPos.Y && checkedPos.X == prevPos.X {
			continue
		}
		if value.Char == models.PathKey || value.Char == models.StartKey {
			edges = append(edges, checkedPos)
		} else if value.Char == models.EndKey {
			edges = append(edges, checkedPos)
		}
	}
	return edges
}

// GenerateGraph is the function that will grenerate our graph using map data
func GenerateGraph(dataMap *models.DataMap) *models.ItemGraph {
	var childPos, prevChildPos models.Coords
	var tmpChildPosEdges, nodePosEdges []models.Coords
	pos := dataMap.GetStart()
	// Init queue
	q := models.NodeQueue{}
	queue := q.New()
	queue.Enqueue(models.Node{Value: *pos})
	// Init graph
	graph := models.ItemGraph{}
	graph.SetStart(&models.Node{Value: *pos})
	for !queue.IsEmpty() {
		nodePos := queue.Front().Value.(models.Coords)
		nodePosEdges = checkEdges(dataMap, models.Coords{Y: -1, X: -1}, nodePos)
		dataMap.AddEdges(nodePos, nodePosEdges)
		// fmt.Printf("Node: %v <-> %v\n", nodePos, nodePosEdges)
		dataMap.PrintMap(nodePos)
		for _, edge := range nodePosEdges {
			if dataMap.GetData(edge).Visited {
				// pretty.Println("Already visited:", edge)
				continue
			}
			prevChildPos = nodePos
			childPos = edge
			tmpChildPosEdges = checkEdges(dataMap, nodePos, edge)
			// fmt.Println("Edge:", edge)
			// dataMap.PrintMap(edge)
			for {
				dataMap.PrintMap(childPos)
				if len(tmpChildPosEdges) > 1 {
					if len(dataMap.GetEdges(childPos)) == 0 {
						graph.HandleNewEdge(&models.Node{Value: nodePos}, &models.Node{Value: childPos})
						if !queue.IsIn(models.Node{Value: childPos}) {
							queue.Enqueue(models.Node{Value: childPos})
						}
					} else {
						if !dataMap.EdgesAllVisited(childPos) {
							// log.Fatal()
						}
					}
					break
				} else if len(tmpChildPosEdges) == 0 {
					graph.HandleNewEdge(&models.Node{Value: nodePos}, &models.Node{Value: childPos})
					// End branch
					dataMap.SetDataVisited(childPos, true)
					break
				}
				// Start node reached
				if models.CompareCoords(childPos, *dataMap.GetStart()) {
					graph.HandleNewEdge(&models.Node{Value: nodePos}, &models.Node{Value: childPos})
					break
				}
				dataMap.SetDataVisited(childPos, true)
				prevChildPos = childPos
				childPos = tmpChildPosEdges[0]
				tmpChildPosEdges = checkEdges(dataMap, prevChildPos, childPos)
			}
		}
		queue.Dequeue()
	}
	dataMap.PrintMap()
	return &graph
}
