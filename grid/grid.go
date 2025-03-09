package grid

import (
	"game_of_life/math"
	"math/rand/v2"
)

const DEAD, ALIVE int = 0, 1

type Grid struct {
	Cells []int
	Size  math.Vector2DInt32
}

func NewGrid(x int32, y int32) *Grid {
	var g = Grid{}
	g.Cells = make([]int, x*y)
	g.Size.X = x
	g.Size.Y = y
	return &g
}

func (g Grid) Randomize() {
	for i := range g.Cells {
		g.Cells[i] = rand.Int() % 2
	}
}

func (g *Grid) Update() {
	var length = len(g.Cells)
	newCells := make([]int, length)

	for index, state := range g.Cells {
		indices := getNeighboringIndices(index, int(g.Size.X), int(g.Size.Y))
		neighborsCount := countAliveCells(g.Cells, indices)
		newCells[index] = getStateFromRules(state, neighborsCount)
	}

	g.Cells = newCells
}

func getNeighboringIndices(index int, gridX int, gridY int) []int {
	const FROM, TO int = -1, 1
	var neighboringIndices []int

	length := gridX * gridY
	for relY := FROM; relY <= TO; relY++ {
		for relX := FROM; relX <= TO; relX++ {
			if relX == 0 && relY == 0 {
				continue
			}
			neighborIndex := index + relX + relY*gridX
			neighborIndex = math.Mod(neighborIndex, length)
			neighboringIndices = append(neighboringIndices, neighborIndex)
		}
	}
	return neighboringIndices
}

func countAliveCells(cells []int, indices []int) int {
	var sum int = 0
	for _, index := range indices { // ! Assumes state == ALIVE, change if the logic gets edited
		sum += cells[index]
	}
	return sum
}

func getStateFromRules(state int, neighborCount int) int {
	if state == ALIVE {
		if neighborCount == 2 || neighborCount == 3 {
			return ALIVE
		} else {
			return DEAD
		}
	} else {
		if neighborCount == 3 {
			return ALIVE
		}
	}
	return state
}
