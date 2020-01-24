package manageMap

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/kr/pretty"
)

// Exec execute map
func Exec() {
	file, err := os.Open("/Users/vomnes/Documents/programming/graph-algorithms-go/input/maze.map")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mapData := DataMap{}
	m := mapData.New()
	y, x := 0, 0
	for scanner.Scan() {
		x = 0
		m.AllocNewY(y)
		strings.Map(func(r rune) rune {
			m.SetData(y, x, r)
			if r == StarKey {
				m.SetStart(y, x)
			}
			if r == EndKey {
				m.SetEnd(y, x)
			}
			x++
			return r
		}, scanner.Text())
		y++
	}
	pretty.Print(m)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
