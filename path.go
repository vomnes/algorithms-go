package main

import (
	"container/list"
	"fmt"
)

type path struct {
	list         *list.List
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
func (d *DataPath) InsertNewPath(item string) *list.List {
	d.listPath.PushBack(
		path{
			list:         list.New(),
			lastItemName: item,
		},
	)
	newPath := d.listPath.Front().Value.(path)
	newPath.list.PushBack(item)
	return newPath.list
}

// CopyPathAddNewItem allows to duplicate a list and append a new item
// to this list
func (d *DataPath) CopyPathAddNewItem(toCopy list.List, newItem string) {
	toCopy.PushBack(newItem)
	d.listPath.PushBack(
		path{
			list:         &toCopy,
			lastItemName: newItem,
		},
	)
	d.Print()
	fmt.Println("->")
}

// RemovePath allows to remove a given path from the path list
func (d *DataPath) RemovePath(toRemove list.List) {
	for e := d.listPath.Front(); e != nil; e = e.Next() {
		if e.Value.(path).list == &toRemove {
			d.listPath.Remove(e)
			break
		}
	}
}

// Print allows to print the path lists
func (d *DataPath) Print() {
	for e := d.listPath.Front(); e != nil; e = e.Next() {
		currentItem := e.Value.(path)
		for item := currentItem.list.Front(); item != nil; item = item.Next() {
			fmt.Print(item.Value, " ")
		}
		fmt.Print("\n")
	}
}
