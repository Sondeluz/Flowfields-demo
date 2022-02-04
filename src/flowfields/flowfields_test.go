package flowfields

import (
	"log"
	"testing"
)

func canGetToObjective(f Flowfield, startPos XYPosition, maxSteps int) bool {
	if startPos.X < 0 || startPos.Y < 0 || startPos.X > GRID_WIDTH || startPos.Y > GRID_HEIGHT {
		panic("Invalid grid position") // This cannot be triggered by the user
	} else if f.grid[startPos.Y][startPos.X].isObstacle { // It makes no sense to start from inside an obstacle
		return true
	}

	currentPos := startPos
	steps := 0

	for !currentPos.equals(f.objective) && steps < maxSteps {
		currentPos = currentPos.advance(f.grid[currentPos.Y][currentPos.X].vector)

		steps += 1
	}

	return steps < maxSteps
}

func TestCanReachObjFromEverywhere(t *testing.T) {
	f := newRandomFlowFieldWithoutObstacles()
	obj := f.getObjective()

	for y := 0; y < GRID_HEIGHT; y++ {
		for x := 0; x < GRID_WIDTH; x++ {
			if !canGetToObjective(f, XYPosition{X: x, Y: y}, 1000) {
				log.Println("Couldn't reach objective", obj, "from (", x, ",", y, ")")
				t.FailNow()
			} else {
				log.Println("Reached", obj, "from (", x, ",", y, ")")
			}
		}
	}

}
