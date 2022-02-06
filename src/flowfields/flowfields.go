package flowfields

import (
	"math"
	//"log"
	//"fmt"
	"math/rand"
	"time"
)

const (
	GRID_HEIGHT = 20
	GRID_WIDTH  = 20
)

type GridPosition struct {
	vector     XYVector
	isObstacle bool
}

type Flowfield struct {
	grid      [][]GridPosition
	objective XYPosition
}

func newRandomFlowFieldWithoutObstacles() Flowfield {
	rand.Seed(time.Now().UnixNano())

	objective := XYPosition{X: rand.Intn(GRID_WIDTH), Y: rand.Intn(GRID_HEIGHT)}

	return newFlowfield(objective, []XYPosition{})
}

func newRandomFlowFieldWithObstacles() Flowfield {
	rand.Seed(time.Now().UnixNano())

	objective := XYPosition{X: rand.Intn(999), Y: rand.Intn(999)}

	obstacles := make([]XYPosition, 0)
	for i := 0; i < 50; i++ {
		obstacles = append(obstacles, XYPosition{X: rand.Intn(999), Y: rand.Intn(999)})
	}

	return newFlowfield(objective, obstacles)
}

func newFlowfield(objective XYPosition, obstacles []XYPosition) (f Flowfield) {
	if objective.X < 0 || objective.Y < 0 || objective.X > GRID_WIDTH || objective.Y > GRID_HEIGHT {
		panic("Invalid objective position") // This cannot be triggered by the user
	}

	f.objective = objective

	// Initialize the grid
	f.grid = make([][]GridPosition, GRID_HEIGHT)
	for y := range f.grid {
		f.grid[y] = make([]GridPosition, GRID_WIDTH) // All GridPositions will be {(0, 0), false} by default
	}

	for _, pos := range obstacles { // Set up obstacles
		f.grid[pos.Y][pos.X].isObstacle = true
	}

	// For each grid square
	for y := range f.grid {
		for x := range f.grid[y] {
			if !f.grid[y][x].isObstacle && !(y == objective.Y) && !(x == objective.X) {
				// Get the neighbour with the lowest distance to the objective, using
				// Manhattan distances
				closestNeighbour := f.getClosestNeighbour(f.getNeighbours(x, y))

				if closestNeighbour.isValid() {
					f.grid[y][x].vector = newXYVector(closestNeighbour.X-x, closestNeighbour.Y-y)
				} else {
					panic("Found an invalid closest neighbour!") // Same
				}
			} else {
                f.grid[y][x].vector = newXYVector(0, 0)
            }
		}
	}

	return f
}

func (f Flowfield) getObjective() XYPosition {
	return f.objective
}

// From https://github.com/yucchiy/minesweeper/blob/c02cede38aa6b046eca399b70b9f6d3e081a11bb/field.go#L130-L150
func (f Flowfield) getNeighbours(x int, y int) (neighbourList []XYPosition) {
	if x < 0 || y < 0 || x > GRID_WIDTH || y > GRID_HEIGHT {
		panic("Invalid grid position") // This cannot be triggered by the user
	}

	neighbourList = make([]XYPosition, 0)

	// For every possible position around the given coordinates
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {

			if dy == 0 && dx == 0 || f.grid[y][x].isObstacle {
				continue
			}

			// Potential neighbouring position
			nX := x + dx
			nY := y + dy

			// If the position is inside the grid
			if nX >= 0 && nY >= 0 && nX < GRID_WIDTH && nY < GRID_HEIGHT {
				// Only add a diagonal neighbor if there aren't any obstacles between them
				if nX == x+1 && nY == y+1 { // bottom-right corner
					if !f.grid[y][x+1].isObstacle && !f.grid[y+1][x].isObstacle {
						neighbourList = append(neighbourList, XYPosition{X: nX, Y: nY})
					}
				} else if nX == x+1 && nY == y-1 { // upper-right corner
					if !f.grid[y][x+1].isObstacle && !f.grid[y-1][x].isObstacle {
						neighbourList = append(neighbourList, XYPosition{X: nX, Y: nY})
					}
				} else if nX == x-1 && nY == y+1 { // bottom-left corner
					if !f.grid[y][x-1].isObstacle && !f.grid[y+1][x].isObstacle {
						neighbourList = append(neighbourList, XYPosition{X: nX, Y: nY})
					}
				} else if nX == x-1 && nY == y-1 { // upper-left corner
					if !f.grid[y][x-1].isObstacle && !f.grid[y-1][x].isObstacle {
						neighbourList = append(neighbourList, XYPosition{X: nX, Y: nY})
					}
				} else {
					neighbourList = append(neighbourList, XYPosition{X: nX, Y: nY})
				}
			}
		}
	}

	//log.Println("Neighbours for:(", x, ",", y, "):", neighbourList)

	return neighbourList
}

func (f Flowfield) getClosestNeighbour(neighbourList []XYPosition) (closestNeihgbour XYPosition) {

	closestNeihgbour.setInvalid()

	minDist := math.MaxInt

	for _, pos := range neighbourList {
		dist := abs(pos.X-f.objective.X) + abs(pos.Y-f.objective.Y)

		if dist < minDist {
			minDist = dist
			closestNeihgbour = pos
		}
	}

	return closestNeihgbour
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
