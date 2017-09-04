// Entry point for testing purposes

package main

import (
	"github.com/avallario/gohlmap/mapio"
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

	hlmap.Shift(0, 0, 320)

	mapio.ExportMap(hlmap, "output.map")

	//	maptree.PrintHLMap(hlmap)

	/*
		t := mapio.Token{-1, "Token init"}
		for t.Kind != mapio.EOF {
			t, err = scanner.Scan()
			check(err)
			fmt.Printf("%v\t<%v>\n", mapio.KindName(t.Kind), t.Spelling)
		}
	*/
}
