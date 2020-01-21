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
func (d *DataPath) CopyPathAddNewItem(toCopy []string, newItem string) {
	toCopy = append(toCopy, newItem)
	d.listPath.PushBack(
		path{
			list:         toCopy,
			lastItemName: newItem,
		},
	)
}

// RemovePath allows to remove a given path from the path list
func (d *DataPath) RemovePath(toRemove []string) {
	for e := d.listPath.Front(); e != nil; e = e.Next() {
		if arraysEqual(e.Value.(path).list, toRemove) {
			d.listPath.Remove(e)
			break
		}
	}
}

// Print allows to print the path lists
func (d *DataPath) Print() {
	for e := d.listPath.Front(); e != nil; e = e.Next() {
		currentItem := e.Value.(path)
		fmt.Println(currentItem.list)
	}
}
