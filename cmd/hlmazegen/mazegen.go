package main

import (
	"github.com/avallario/gohlmap/cmd/hlmazegen/generator"
	"github.com/avallario/gohlmap/mapio"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	hlmap := generator.GenerateMap(generator.GenerateMaze(10, 10))
	mapio.ExportMap(hlmap, "output.map")
}
