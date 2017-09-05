package mapio

import (
	"github.com/avallario/gohlmap/maptree"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ImportMap(filename string) *maptree.HLMap {
	f, err := os.Open(filename)
	defer f.Close()
	check(err)

	scanner := NewScanner(f)
	parser := NewParser(scanner)

	return parser.Parse()
}
