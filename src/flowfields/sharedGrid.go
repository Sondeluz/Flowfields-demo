package flowfields

import (
	"sync"
)

const (
	UNOCCUPIED = -1
)

// A position on the shared map, with
// an optional agent ID who may be going
// to occupy it in its next move
type SharedGridPosition struct {
	agentID int
	mtx     sync.Mutex // mutex for this position
}

// Shared grid between all agents
type SharedGrid struct {
	grid [][]SharedGridPosition
}

func newSharedGrid() *SharedGrid {
	grid := make([][]SharedGridPosition, GRID_HEIGHT)
	for y := range grid {
		grid[y] = make([]SharedGridPosition, GRID_WIDTH)
		for col := range grid[y] {
			grid[y][col].agentID = UNOCCUPIED
		}
	}

	return &SharedGrid{grid}
}

// Attempts to occupy a position in the shared grid, returning true if success,
// false otherwise
func (s *SharedGrid) attemptToOccupy(pos XYPosition, agent Agent) bool {
	if pos.X < 0 || pos.Y < 0 || pos.X > GRID_WIDTH || pos.Y > GRID_HEIGHT {
		panic("Invalid shared grid position")
	}

	s.grid[pos.Y][pos.X].mtx.Lock()
	defer s.grid[pos.Y][pos.X].mtx.Unlock()

	if s.grid[pos.Y][pos.X].agentID == UNOCCUPIED {
		s.grid[pos.Y][pos.X].agentID = agent.id // Occupy it
		return true
	} else {
		return false
	}
}

// Sets a position in the shared grid as unoccupied
func (s *SharedGrid) free(pos XYPosition) {
	if pos.X < 0 || pos.Y < 0 || pos.X > GRID_WIDTH || pos.Y > GRID_HEIGHT {
		panic("Invalid shared grid position")
	}

	s.grid[pos.Y][pos.X].mtx.Lock()
	defer s.grid[pos.Y][pos.X].mtx.Unlock()

	s.grid[pos.Y][pos.X].agentID = UNOCCUPIED
}
