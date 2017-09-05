package generator

import (
	"math/rand"
)

const (
	NORTH      = iota
	SOUTH      = iota
	EAST       = iota
	WEST       = iota
	SPAWN_RATE = 0.4
)

type MazeCell struct {
	N, S, E, W bool
	visited    bool
}

type pathFinder struct {
	x, y int
	dead bool
}

func GenerateMaze(w, h int) [][]*MazeCell {
	maze := make([][]*MazeCell, w)
	for i := range maze {
		maze[i] = make([]*MazeCell, h)
		for j := range maze[i] {
			maze[i][j] = new(MazeCell)
			maze[i][j].N = true
			maze[i][j].S = true
			maze[i][j].E = true
			maze[i][j].W = true
		}
	}
	maze[0][0].visited = true

	var finders []*pathFinder
	finders = append(finders, new(pathFinder))

	for len(finders) > 0 {
		var new_finders []*pathFinder

		for i, f := range finders {
			new_finder := f.act(maze)
			if new_finder != nil {
				new_finders = append(new_finders, new_finder)
			}
		}

		for i, f := range finders {
			if f.dead {
				copy(finders[i:], finders[i+1:])
				finders[len(finders)-1] = nil
				finders = finders[:len(finders)-1]
			}
		}

		finders = append(new_finders, finders...)
	}

	return maze
}

// Returns a new pathfinder if one was spawned, true if the pathfinder should be removed
func (p *pathFinder) act(maze [][]*MazeCell) *pathFinder {
	var tried_north, tried_south, tried_east, tried_west bool
	var new_finder *pathFinder
	var dir int
	max_x := len(maze) - 1
	max_y := len(maze[0]) - 1

	for !(tried_north && tried_south && tried_east && tried_west) {
		dir = rand.Intn(4)
		acted := false

		switch dir {
		case NORTH:
			if !tried_north {
				tried_north = true
				if p.y < max_y && !maze[p.x][p.y+1].visited {
					maze[p.x][p.y].N = false
					maze[p.x][p.y+1].S = false
					maze[p.x][p.y+1].visited
					p.y++
					acted = true
				}
			}
		case SOUTH:
			if !tried_south {
				tried_south = true
				if p.y > 0 && !maze[p.x][p.y-1].visited {
					maze[p.x][p.y].S = false
					maze[p.x][p.y-1].N = false
					maze[p.x][p.y-1].visited = true
					p.y--
					acted = true
				}
			}
		case EAST:
			if !tried_east {
				tried_east = true
				if p.x < max_x && !maze[p.x+1][p.y].visited {
					maze[p.x][p.y].E = false
					maze[p.x+1][p.y].W = false
					maze[p.x+1][p.y].visited = true
					p.x++
					acted = true
				}
			}
		case WEST:
			if !tried_west {
				tried_west = true
				if p.x > 0 && !maze[p.x-1][p.y].visited {
					maze[p.x][p.y].W = false
					maze[p.x-1][p.y].E = false
					maze[p.x-1][p.y].visited = true
					p.x--
					acted = true
				}
			}
		}

		if acted {
			if rand.Float64() <= SPAWN_RATE {
				new_finder = new(pathFinder)
				new_finder.x = p.x
				new_finder.y = p.y
			}
			return new_finder
		}
	}

	p.dead = true
	return nil
}
