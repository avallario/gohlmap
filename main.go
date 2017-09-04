// Entry point for testing purposes

package main

import (
	"github.com/avallario/gohlmap/mapio"
	"github.com/avallario/gohlmap/maptree"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("testinput.map")
	defer f.Close()
	check(err)

	scanner := mapio.NewScanner(f)
	parser := mapio.NewParser(scanner)

	hlmap := parser.Parse()

	entity, err2 := hlmap.FindEntityNamed("rail")
	check(err2)

	err3 := hlmap.MoveEntityToWorld(entity)
	check(err3)

	//	hlmap.Shift(0, 0, 320)

	mapio.ExportMap(hlmap, "output.map")

	maptree.PrintHLMap(hlmap)
}
