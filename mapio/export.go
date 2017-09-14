package mapio

import (
	"bufio"
	"fmt"
	"github.com/avallario/gohlmap/maptree"
	"os"
)

func ExportMap(hlmap *maptree.HLMap, filename string) {
	f, err := os.Create(filename)
	defer f.Close()
	check(err)
	output := bufio.NewWriter(f)

	for _, entity := range hlmap.Entitylist {
		exportEntity(entity, output)
	}

	output.Flush()
}

func exportEntity(entity *maptree.Entity, output *bufio.Writer) {
	var err error

	_, err = output.WriteString("{\n")
	check(err)

	for k, v := range entity.Properties {
		_, err = output.WriteString("\"" + k + "\" \"" + v + "\"\n")
		check(err)
	}

	for _, brush := range entity.Brushlist {
		exportBrush(brush, output)
	}

	_, err = output.WriteString("}\n")
	check(err)
}

func exportBrush(brush *maptree.Brush, output *bufio.Writer) {
	var err error

	_, err = output.WriteString("\t{\n")
	check(err)

	for _, face := range brush.Facelist {
		face_string := fmt.Sprintf("\t\t( %v %v %v ) ( %v %v %v ) ( %v %v %v ) %v [ %v %v %v %v ] [ %v %v %v %v ] %v %v %v\n", face.X1, face.Y1, face.Z1, face.X2, face.Y2, face.Z2, face.X3, face.Y3, face.Z3, face.Texname, face.TX1.FloatString(6), face.TY1.FloatString(6), face.TZ1.FloatString(6), face.TOffset1, face.TX2.FloatString(6), face.TY2.FloatString(6), face.TZ2.FloatString(6), face.TOffset2, face.Rot.FloatString(2), face.ScaleX.FloatString(2), face.ScaleY.FloatString(2))
		_, err = output.WriteString(face_string)
		check(err)
	}

	_, err = output.WriteString("\t}\n")
	check(err)
}
