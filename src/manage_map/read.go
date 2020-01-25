package manageMap

import (
	"bufio"
	"log"
	"os"
	"strings"

	"../models"
)

// StoreInputMap read the path file and store the data in an structure
func storeInputMap(path string) *models.DataMap {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mapData := models.DataMap{}
	m := mapData.New()
	y, x := 0, 0
	for scanner.Scan() {
		x = 0
		m.AllocNewY(y)
		strings.Map(func(r rune) rune {
			m.SetDataChar(y, x, r)
			if r == models.StartKey {
				m.SetStart(y, x)
			}
			if r == models.EndKey {
				m.SetEnd(y, x)
			}
			x++
			return r
		}, scanner.Text())
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m
}
