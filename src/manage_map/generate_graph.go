package manageMap

import (
	"fmt"

	"../models"
)

func checkEdges(m *models.DataMap, x, y int) {

}

// GenerateGraph is the function that will grenerate our graph using map data
func GenerateGraph(dataMap *models.DataMap) {
	graph := models.ItemGraph{}
	g := graph.String()

	fmt.Print(g)
}
