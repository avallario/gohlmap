package generator

import (
	"github.com/avallario/gohlmap/mapio"
	"github.com/avallario/gohlmap/maptree"
	"strconv"
)

const (
	CELL_SIZE      = 256
	SOUTH_BOUNDARY = -2560
	WEST_BOUNDARY  = -2560
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GenerateMap(maze [][]*MazeCell) *maptree.HLMap {
	hlmap := mapio.ImportMap("empty.map")
	hlmap.FindWorldspawn()
	cell := mapio.ImportMap("cell.map")

	for x := range maze {
		for y := range maze[0] {
			new_cell := cell.Copy()
			new_cell.FindWorldspawn()

			if !maze[x][y].N {
				e, err := new_cell.FindEntityNamed("north")
				check(err)
				new_cell.Entitylist = removeEnt(new_cell.Entitylist, e)
			}
			if !maze[x][y].S {
				e, err := new_cell.FindEntityNamed("south")
				check(err)
				new_cell.Entitylist = removeEnt(new_cell.Entitylist, e)
			}
			if !maze[x][y].E {
				e, err := new_cell.FindEntityNamed("east")
				check(err)
				new_cell.Entitylist = removeEnt(new_cell.Entitylist, e)
			}
			if !maze[x][y].W {
				e, err := new_cell.FindEntityNamed("west")
				check(err)
				new_cell.Entitylist = removeEnt(new_cell.Entitylist, e)
			}

			ents_to_remove := make([]*maptree.Entity, len(new_cell.Entitylist))
			copy(ents_to_remove, new_cell.Entitylist)

			for _, entity := range ents_to_remove {
				if entity != new_cell.Worldspawn() {
					new_cell.MoveEntityToWorld(entity)
				}
			}

			new_cell.Shift(WEST_BOUNDARY+CELL_SIZE*x, SOUTH_BOUNDARY+CELL_SIZE*y, 0)

			hlmap.Worldspawn().Brushlist = append(hlmap.Worldspawn().Brushlist, new_cell.Worldspawn().Brushlist...)
		}
	}

	player_start := new(maptree.Entity)
	player_start.Properties = make(map[string]string)
	player_start.Properties["classname"] = "info_player_start"
	player_start.Properties["angles"] = "0 0 0"
	player_start.Properties["origin"] = strconv.Itoa(WEST_BOUNDARY+CELL_SIZE/2) + " " + strconv.Itoa(SOUTH_BOUNDARY+CELL_SIZE/2) + " 40"
	hlmap.Entitylist = append(hlmap.Entitylist, player_start)

	return hlmap
}

func removeEnt(entities []*maptree.Entity, e *maptree.Entity) []*maptree.Entity {
	index := -1
	for i, entity := range entities {
		if entity == e {
			index = i
		}
	}

	if index == -1 {
		return entities
	}

	copy(entities[index:], entities[index+1:])
	entities[len(entities)-1] = nil
	entities = entities[:len(entities)-1]

	return entities
}
