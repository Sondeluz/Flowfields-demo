package flowfields

import (
	"math/rand"
	"time"
)

type Agent struct {
	id         int
	position   XYPosition
	flowfield  Flowfield
	sharedGrid *SharedGrid // Shared flowfield between all agents
}

func NewAgent(id int, startPos XYPosition, flowfield Flowfield, sharedGrid *SharedGrid) (a Agent) {
	a.id = id
	a.position = startPos
	a.flowfield = flowfield
	a.sharedGrid = sharedGrid

	return a
}

// Move to the next position in the flowfield, or elsewhere (or nowhere) if said next
// position is taken by another agent
func (a *Agent) move() {
	newPos := a.position.advance(a.flowfield.grid[a.position.Y][a.position.X].vector)

	if !a.sharedGrid.attemptToOccupy(newPos, *a) {
		a.moveElsewhere(newPos)
	} else {
		a.sharedGrid.free(a.position)
		a.position = newPos
	}
}

// Attempt to move elsewhere given an old desired position which is occupied
func (a *Agent) moveElsewhere(oldPos XYPosition) {
	neighbours := a.flowfield.getNeighbours(a.position.X, a.position.Y)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(neighbours), func(i, j int) { neighbours[i], neighbours[j] = neighbours[j], neighbours[i] })

	// Attempt to occupy at random any of the valid
	// neighbours which are free on the shared grid
	for _, n := range neighbours {
		if !n.equals(oldPos) && a.sharedGrid.attemptToOccupy(n, *a) {
			a.sharedGrid.free(a.position)
			a.position = n
			return
		}
	}

	// If no success, it will remain still and try again later in its next turn
}

func (a Agent) isInObjective() bool {
	return a.position.equals(a.flowfield.objective)
}
