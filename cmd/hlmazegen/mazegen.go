package main

import (
	"github.com/avallario/gohlmap/cmd/hlmazegen/generator"
	"github.com/avallario/gohlmap/mapio"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//TODO: Take command line arguments for setting output_filename_root and num_mazes
	output_filename_root := "randmaze"
	num_mazes := 1

	for i := 0; i < num_mazes; i++ {
		output_filename := output_filename_root + strconv.Itoa(i)

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
}
