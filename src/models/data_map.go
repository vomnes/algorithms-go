package models

import "sync"

const (
	// StartKey is the char key to indicate the map start
	StartKey = 'S'
	// EndKey is the char key to indicate the map end
	EndKey = 'E'
	// WallKey is the char key to indicate a wall position
	WallKey = '#'
	// PathKey is the char key to indicate a path position
	PathKey = '.'
)

// Coords store the X, Y integers used as coordinates
type Coords struct {
	X, Y int
}

// LinkEdge is the data structure of a edge of a link
type LinkEdge struct {
	Position Coords
	Explored bool
}

// Link is the data structure of the map item
type Link struct {
	Char       rune
	ChildEdges []LinkEdge
}

// DataMap store the map data
type DataMap struct {
	// data -> y x char
	data  map[int]map[int]Link
	start Coords
	end   Coords
	lock  sync.RWMutex
}

// New creates a new data map
func (s *DataMap) New() *DataMap {
	s.lock.Lock()
	s.data = make(map[int]map[int]Link)
	s.start = Coords{}
	s.end = Coords{}
	s.lock.Unlock()
	return s
}

// AllocNewY allocates a new y array range in the map
func (s *DataMap) AllocNewY(y int) {
	s.lock.Lock()
	s.data[y] = make(map[int]Link)
	s.lock.Unlock()
}

// SetDataChar set the rune at the y x position
func (s *DataMap) SetDataChar(y, x int, value rune) {
	s.lock.Lock()
	tmp := s.data[y][x]
	tmp.Char = value
	s.data[y][x] = tmp
	s.lock.Unlock()
}

// SetStart set the x, y data of the start node
func (s *DataMap) SetStart(y, x int) {
	s.lock.Lock()
	s.start = Coords{
		Y: y,
		X: x,
	}
	s.lock.Unlock()
}

// SetEnd set the x, y data of the end node
func (s *DataMap) SetEnd(y, x int) {
	s.lock.Lock()
	s.end = Coords{
		Y: y,
		X: x,
	}
	s.lock.Unlock()
}

// AddEdges appends given edges to a given position in the map
func (s *DataMap) AddEdges(pos Coords, edges []Coords) {
	s.lock.Lock()
	tmp := s.data[pos.Y][pos.Y]
	for _, edge := range edges {
		tmp.ChildEdges = append(
			tmp.ChildEdges,
			LinkEdge{
				Position: edge,
				Explored: false,
			},
		)
	}
	s.data[pos.Y][pos.Y] = tmp
	s.lock.Unlock()
}

// GetEdges return the linkEdge array of a given position
func (s *DataMap) GetEdges(pos Coords) []LinkEdge {
	return s.data[pos.Y][pos.Y].ChildEdges
}

// EdgesAllVisited return true if all the childs of a position are visited
func (s *DataMap) EdgesAllVisited(pos Coords) bool {
	edges := s.data[pos.Y][pos.Y].ChildEdges
	for _, edge := range edges {
		if edge.Explored == false {
			return false
		}
	}
	return true
}

// GetDataChar return the rune value of the data at the y, x coordinates
func (s *DataMap) GetDataChar(y, x int) rune {
	return s.data[y][x].Char
}

// GetStart return the y, x coordinates of the start positon
func (s *DataMap) GetStart() *Coords {
	return &s.start
}
