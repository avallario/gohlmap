package maptree

import (
	"fmt"
	"math/big"
)

type treeError struct {
	msg string
}

func (t treeError) Error() string {
	return "Tree operation error: " + t.msg
}

// ---------- Face class definition ----------

type Face struct {
	X1, Y1, Z1         int
	X2, Y2, Z2         int
	X3, Y3, Z3         int
	Texname            string
	TOffset1, TOffset2 int
	TX1, TY1, TZ1      *big.Rat // Precision 6
	TX2, TY2, TZ2      *big.Rat // Precision 6
	Rot                *big.Rat // Precision 2
	ScaleX, ScaleY     *big.Rat // Precision 2
}

func (f *Face) Shift(dx, dy, dz int) {
	f.X1 += dx
	f.X2 += dx
	f.X3 += dx
	f.Y1 += dy
	f.Y2 += dy
	f.Y3 += dy
	f.Z1 += dz
	f.Z2 += dz
	f.Z3 += dz
}

func (f *Face) Copy() *Face {
	new_face := new(Face)
	*new_face = *f
	return new_face
}

// ---------- Brush class definition ----------

type Brush struct {
	Facelist []*Face
}

func (b *Brush) Shift(dx, dy, dz int) {
	for _, f := range b.Facelist {
		f.Shift(dx, dy, dz)
	}
}

func (b *Brush) Copy() *Brush {
	new_brush := new(Brush)

	for _, face := range b.Facelist {
		new_brush.Facelist = append(new_brush.Facelist, face.Copy())
	}

	return new_brush
}

// ---------- Entity class definition ----------

type Entity struct {
	Brushlist  []*Brush
	Properties map[string]string
}

func (e *Entity) Shift(dx, dy, dz int) {
	for _, b := range e.Brushlist {
		b.Shift(dx, dy, dz)
	}
}

func (e *Entity) Copy() *Entity {
	new_entity := new(Entity)
	new_entity.Properties = make(map[string]string)

	for k, v := range e.Properties {
		new_entity.Properties[k] = v
	}

	for _, brush := range e.Brushlist {
		new_entity.Brushlist = append(new_entity.Brushlist, brush.Copy())
	}

	return new_entity
}

// ---------- HLMap class definition ----------

type HLMap struct {
	Entitylist []*Entity
	worldspawn *Entity
}

func (q *HLMap) Shift(dx, dy, dz int) {
	for _, e := range q.Entitylist {
		e.Shift(dx, dy, dz)
	}
}

func (q *HLMap) MoveEntityToWorld(e *Entity) error {
	ws := q.Worldspawn()

	if ws == nil {
		return treeError{"HLMap has no identified worldspawn (call .FindWorldspawn())"}
	}

	found_entity := false
	entity_index := -1

	for i, entity := range q.Entitylist {
		if e == entity {
			found_entity = true
			entity_index = i
			break
		}
	}

	if found_entity {
		ws.Brushlist = append(ws.Brushlist, e.Brushlist...)

		copy(q.Entitylist[entity_index:], q.Entitylist[entity_index+1:])
		q.Entitylist[len(q.Entitylist)-1] = nil
		q.Entitylist = q.Entitylist[:len(q.Entitylist)-1]

		return nil
	} else {
		return treeError{"Could not find specified entity in map object"}
	}
}

func (q *HLMap) FindEntityNamed(name string) (*Entity, error) {
	for _, entity := range q.Entitylist {
		for k, v := range entity.Properties {
			if k == "targetname" && v == name {
				return entity, nil
			}
		}
	}
	return nil, treeError{"Could not find entity with specified name in map object"}
}

func (q *HLMap) Copy() *HLMap {
	new_hlmap := new(HLMap)

	for _, entity := range q.Entitylist {
		new_hlmap.Entitylist = append(new_hlmap.Entitylist, entity.Copy())
	}

	return new_hlmap
}

func (q *HLMap) FindWorldspawn() error {
	found_worldspawn := false

	for _, entity := range q.Entitylist {
		for k, v := range entity.Properties {
			if k == "classname" && v == "worldspawn" {
				found_worldspawn = true
				q.worldspawn = entity
				break
			}
		}

		if found_worldspawn {
			break
		}
	}

	if found_worldspawn {
		return nil
	} else {
		return treeError{"Unable to find worldspawn"}
	}
}

func (q *HLMap) Worldspawn() *Entity {
	return q.worldspawn
}

func PrintHLMap(q *HLMap) {
	for _, entity := range q.Entitylist {
		fmt.Printf("ENTITY @ %p\n", entity)
		for k, v := range entity.Properties {
			fmt.Printf("\t%v : %v\n", k, v)
		}
		for _, brush := range entity.Brushlist {
			fmt.Printf("\tBRUSH @ %p\n", brush)
			for _, face := range brush.Facelist {
				/*
					fmt.Println("\t\tFACE")
					fmt.Printf("\t\t\tX1\t%v\n", face.X1)
					fmt.Printf("\t\t\tY1\t%v\n", face.Y1)
					fmt.Printf("\t\t\tZ1\t%v\n", face.Z1)
					fmt.Printf("\t\t\tX2\t%v\n", face.X2)
					fmt.Printf("\t\t\tY2\t%v\n", face.Y2)
					fmt.Printf("\t\t\tZ2\t%v\n", face.Z2)
					fmt.Printf("\t\t\tX3\t%v\n", face.X3)
					fmt.Printf("\t\t\tY3\t%v\n", face.Y3)
					fmt.Printf("\t\t\tZ3\t%v\n", face.Z3)
					fmt.Printf("\t\t\tTexname\t%v\n", face.Texname)
					fmt.Printf("\t\t\tTX1\t%v\n", face.TX1.FloatString(6))
					fmt.Printf("\t\t\tTY1\t%v\n", face.TY1.FloatString(6))
					fmt.Printf("\t\t\tTZ1\t%v\n", face.TZ1.FloatString(6))
					fmt.Printf("\t\t\tTOffset1\t%v\n", face.TOffset1)
					fmt.Printf("\t\t\tTX2\t%v\n", face.TX2.FloatString(6))
					fmt.Printf("\t\t\tTY2\t%v\n", face.TY2.FloatString(6))
					fmt.Printf("\t\t\tTZ2\t%v\n", face.TZ2.FloatString(6))
					fmt.Printf("\t\t\tTOffset2\t%v\n", face.TOffset2)
					fmt.Printf("\t\t\tRot\t%v\n", face.Rot.FloatString(6))
					fmt.Printf("\t\t\tScaleX\t%v\n", face.ScaleX.FloatString(6))
					fmt.Printf("\t\t\tScaleY\t%v\n", face.ScaleY.FloatString(6))
				*/

				fmt.Printf("\t\t\t( %v %v %v ) ( %v %v %v ) ( %v %v %v ) %v [ %v %v %v %v ] [ %v %v %v %v ] %v %v %v\n", face.X1, face.Y1, face.Z1, face.X2, face.Y2, face.Z2, face.X3, face.Y3, face.Z3, face.Texname, face.TX1.FloatString(6), face.TY1.FloatString(6), face.TZ1.FloatString(6), face.TOffset1, face.TX2.FloatString(6), face.TY2.FloatString(6), face.TZ2.FloatString(6), face.TOffset2, face.Rot.FloatString(2), face.ScaleX.FloatString(2), face.ScaleY.FloatString(2))
			}
		}
	}
	ws := q.Worldspawn()
	if ws != nil {
		fmt.Printf("Worldspawn identified as ENTITY @ %p\n", ws)
	}
}
