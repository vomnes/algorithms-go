package algorithms

import (
	"fmt"
	"strings"

	"../models"
)

// Trie contrains the tree
type Trie struct {
	tree *models.ItemGraph
}

func (s *Trie) find(str string) bool {
	var currentNode *models.Node
	var childNodes []*models.Node
	var discoveredStrPart string

	currentNode = s.tree.GetStart()
	for _, char := range str {
		discoveredStrPart += string(char)
		childNodes = s.tree.GetEdges(currentNode)
		for _, child := range childNodes {
			if child.Value == discoveredStrPart {
				currentNode = child
				continue
			}
		}
	}
	// Not found
	if !(currentNode.Value == str) {
		return false
	}
	// Check if the found str is a valid end
	for _, child := range s.tree.GetEdges(currentNode) {
		if child == s.tree.GetEnd() {
			return true
		}
	}
	return false
}

func (s *Trie) insert(str string) {
	var currentNode, newNode *models.Node
	var childNodes []*models.Node
	var discoveredStrPart string

	currentNode = s.tree.GetStart()
	discoveredStrPart = ""
	strings.Map(func(char rune) rune {
		discoveredStrPart += string(char)
		childNodes = s.tree.GetEdges(currentNode)
		for _, child := range childNodes {
			if child.Value == discoveredStrPart {
				currentNode = child
				return char
			}
		}
		newNode = &models.Node{Value: discoveredStrPart}
		s.tree.AddNode(newNode)
		s.tree.AddEdge(currentNode, newNode)
		currentNode = newNode
		return char
	}, str)
	s.tree.AddEdge(s.tree.GetEnd(), currentNode)
}

// TrieExec exec
func TrieExec() {
	trie := &Trie{tree: &models.ItemGraph{}}
	nStart := models.Node{Value: ""}
	nEnd := models.Node{Value: ""}
	trie.tree.AddNode(&nStart)
	trie.tree.AddNode(&nEnd)
	trie.tree.SetStart(&nStart)
	trie.tree.SetEnd(&nEnd)
	trie.insert("Hello")
	trie.insert("Help")
	trie.insert("Bonjour")
	trie.insert("Bon")
	trie.insert("Bonne")
	trie.insert("Bye")

	fmt.Println("Bonjour", trie.find("Bonjour"))
	fmt.Println("Bonjours", trie.find("Bonjours"))
	fmt.Println("Bonjou", trie.find("Bonjou"))
	fmt.Println("Hello", trie.find("Hello"))
	// trie.tree.Print()
}
