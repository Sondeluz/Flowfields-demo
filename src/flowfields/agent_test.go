package flowfields

import (
	"log"
	"testing"
)

const (
	MAX_MOVES = 1000
)

func TestOneAgent(t *testing.T) {
	f := newRandomFlowFieldWithoutObstacles()
	obj := f.getObjective()
	flowfield := newRandomFlowFieldWithoutObstacles()
	sg := newSharedGrid()

	for y := 0; y < GRID_HEIGHT; y++ {
		for x := 0; x < GRID_WIDTH; x++ {

			a := NewAgent(0, XYPosition{X: x, Y: y}, flowfield, sg)
			moves := 0
			for !a.isInObjective() && moves < MAX_MOVES {
				a.move()
			}

			if moves == MAX_MOVES {
				log.Println("Couldn't reach objective", obj, "from (", x, ",", y, ")")
				t.FailNow()
			} else {
				log.Println("Reached", obj, "from (", x, ",", y, ")")
			}

			a.sharedGrid.free(a.position) // Free the objective
		}
	}

}
