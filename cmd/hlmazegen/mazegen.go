package main

import (
	"github.com/avallario/gohlmap/cmd/hlmazegen/generator"
	"github.com/avallario/gohlmap/mapio"
	"math/rand"
	"os/exec"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	output_filename := "maze"

	rand.Seed(time.Now().UnixNano())
	hlmap := generator.GenerateMap(generator.GenerateMaze(20, 20))
	mapio.ExportMap(hlmap, "output.map")

	csg_cmd := exec.Command("hlcsg", "output.map")
	csg_err := csg_cmd.Run()
	check(csg_err)

	bsp_cmd := exec.Command("hlbsp", "output.map")
	bsp_err := bsp_cmd.Run()
	check(bsp_err)

	vis_cmd := exec.Command("hlvis", "output.map")
	vis_err := vis_cmd.Run()
	check(vis_err)

	ren_cmd := exec.Command("cleanup.bat", output_filename+".bsp")
	ren_err := ren_cmd.Run()
	check(ren_err)
}
