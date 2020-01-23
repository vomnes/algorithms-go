package main

import (
	"fmt"
	"sync"
)

type item interface{}

// Node a single node that composes the tree
type Node struct {
	value item
}

// HeuristicItem is the value of the node with heuristic information
type HeuristicItem struct {
	data      string
	heuristic int
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// ItemGraph the Items graph
type ItemGraph struct {
	nodes       []*Node
	edges       map[Node][]*Node
	edgesWeight map[Node][]int
	lock        sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node, weight ...int) {
	g.lock.Lock()
	if g.edges == nil || g.edgesWeight == nil {
		g.edges = make(map[Node][]*Node)
		g.edgesWeight = make(map[Node][]int)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	if weight != nil {
		g.edgesWeight[*n1] = append(g.edgesWeight[*n1], weight[0])
		g.edgesWeight[*n2] = append(g.edgesWeight[*n2], weight[0])
	}
	g.lock.Unlock()
}

// String return the graph links under string format
func (g *ItemGraph) String() string {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	g.lock.RUnlock()
	return s
}

// Print print the graph using string function
func (g *ItemGraph) Print() {
	fmt.Print(g.String())
}

// GetDistance return the weight of a given edge of a node
func (g *ItemGraph) GetDistance(node *Node, edgeIndex int) int {
	return g.edgesWeight[*node][edgeIndex]
}

// GetHeuristicValue return the heuristic value of a given node
func (g *ItemGraph) GetHeuristicValue(node *Node) int {
	for _, elem := range g.nodes {
		if elem == node {
			return node.value.(HeuristicItem).heuristic
		}
	}
	return -1
}
