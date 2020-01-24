package models

import "sync"

const (
	// StarKey is the char key to indicate the map start
	StarKey = 'S'
	// EndKey is the char key to indicate the map end
	EndKey = 'E'
)

type coords struct {
	x, y int
}

// DataMap store the map data
type DataMap struct {
	// data -> y x char
	data  map[int]map[int]rune
	start coords
	end   coords
	lock  sync.RWMutex
}

// New creates a new data map
func (s *DataMap) New() *DataMap {
	s.lock.Lock()
	s.data = make(map[int]map[int]rune)
	s.start = coords{}
	s.end = coords{}
	s.lock.Unlock()
	return s
}

// AllocNewY allocates a new y array range in the map
func (s *DataMap) AllocNewY(y int) {
	s.lock.Lock()
	s.data[y] = make(map[int]rune)
	s.lock.Unlock()
}

// SetData set the rune at the y x position
func (s *DataMap) SetData(y, x int, value rune) {
	s.lock.Lock()
	s.data[y][x] = value
	s.lock.Unlock()
}

// SetStart set the x, y data of the start node
func (s *DataMap) SetStart(y, x int) {
	s.lock.Lock()
	s.start = coords{
		y: y,
		x: x,
	}
	s.lock.Unlock()
}

// SetEnd set the x, y data of the end node
func (s *DataMap) SetEnd(y, x int) {
	s.lock.Lock()
	s.end = coords{
		y: y,
		x: x,
	}
	s.lock.Unlock()
}

// GetData return the rune value of the data at the y, x coordinates
func (s *DataMap) GetData(y, x int) rune {
	return s.data[y][x]
}
