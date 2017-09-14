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

	mapio.ExportMap(hlmap, "output.map")

	maptree.PrintHLMap(hlmap)
}
