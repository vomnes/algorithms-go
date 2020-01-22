package main

import (
	"container/list"
	"fmt"
)

type path struct {
	list         []string
	lastItemName string
}

// DataPath ...
type DataPath struct {
	listPath *list.List
}

// NewListPath ...
func (d *DataPath) NewListPath() *DataPath {
	d.listPath = list.New()
	return d
}

// InsertNewPath ...
func (d *DataPath) InsertNewPath(item string) []string {
	newPath := []string{item}
	d.listPath.PushBack(
		path{
			list:         newPath,
			lastItemName: item,
		},
	)
	return newPath
}

// CopyPathAddNewItem allows to duplicate a list and append a new item
// to this list
func (d *DataPath) CopyPathAddNewItem(toCopy []string, newItem string) []string {
	tmpList := make([]string, len(toCopy))
	copy(tmpList, toCopy)
	tmpList = append(tmpList, newItem)
	d.listPath.PushBack(
		path{
			list:         tmpList,
			lastItemName: newItem,
		},
	)
	return tmpList
}

// RemovePath allows to remove a given path from the path list
func (d *DataPath) RemovePath(toRemove string) {
	for elem := d.listPath.Front(); elem != nil; elem = elem.Next() {
		if elem.Value.(path).lastItemName == toRemove {
			d.listPath.Remove(elem)
			break
		}
	}
}

// Iterate is a fonction that allows to iterate through the path list
// and apply a func for each path
func (d *DataPath) Iterate(f func(path)) {
	for e := d.listPath.Front(); e != nil; e = e.Next() {
		f(e.Value.(path))
	}
}

// Print allows to print the path lists
func (d *DataPath) Print() {
	d.Iterate(func(elem path) {
		fmt.Println("List:", elem.list, "Last Item:", elem.lastItemName)
	})
}

// Growth golang
func (d *DataPath) Growth(currentNodeValue string, edge string) {
	d.Iterate(func(elem path) {
		if elem.lastItemName == currentNodeValue {
			d.CopyPathAddNewItem(elem.list, edge)
		}
	})
}
