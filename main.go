// Entry point for testing purposes

package main

import (
	"fmt"
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
	fmt.Println("OK Good luck test!")

	f, err := os.Open("testinput.map")
	defer f.Close()
	check(err)

	scanner := mapio.NewScanner(f)
	parser := mapio.NewParser(scanner)

	hlmap := parser.Parse()

	maptree.PrintHLMap(hlmap)

	/*
		t := mapio.Token{-1, "Token init"}
		for t.Kind != mapio.EOF {
			t, err = scanner.Scan()
			check(err)
			fmt.Printf("%v\t<%v>\n", mapio.KindName(t.Kind), t.Spelling)
		}
	*/
}
