package manageMap

import "github.com/kr/pretty"

func Exec() {
	m := StoreInputMap("/Users/vomnes/Documents/programming/graph-algorithms-go/input/maze.map")
	pretty.Print(m)
}
